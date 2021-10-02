package gdt

import (
	"strconv"
	"strings"
	"time"
)

// All known fields.
var (
	// alnum, float and num are defined by the specification
	FieldTypeText   FieldType = "alnum"
	FieldTypeFloat  FieldType = "float"
	FieldTypeNumber FieldType = "num"

	// date and time are "virtual" types that are defined as "alnum" but
	// enforce a rather hard set of rules on how the value is structured.
	// Thus, we represent them as dedicated types here to enforce consistent
	// Date (DDMMYYYY) and Time (HHMMSS) handling.
	FieldTypeDate FieldType = "date" // is parsed as time.Time
	FieldTypeTime FieldType = "time" // is parsed as time.Duration so you can use FieldTypeDate (time.Time) .Add(duration).
)

func verifyLength(desc *FieldDesc, data []byte) error {
	if desc.Length != 0 && len(data) != desc.Length {
		return ErrInvalidLength
	}
	if desc.MaxLength != 0 && len(data) >= desc.MaxLength {
		return ErrInvalidLength
	}
	return nil
}

// Register common types
func init() {

	MustRegister(FieldTypeText, func(value []byte, desc *FieldDesc) (interface{}, error) {
		s := string(value)
		if err := verifyLength(desc, value); err != nil {
			return s, err
		}
		return s, nil
	})

	MustRegister(FieldTypeNumber, func(value []byte, desc *FieldDesc) (interface{}, error) {
		if err := verifyLength(desc, value); err != nil {
			return nil, err
		}
		// Numbers must be prefixed with 0 to match the length of the field
		s := strings.TrimLeft(string(value), "0")

		n, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			return nil, err
		}
		return int(n), nil
	})

	MustRegister(FieldTypeFloat, func(value []byte, desc *FieldDesc) (interface{}, error) {
		if err := verifyLength(desc, value); err != nil {
			return nil, err
		}
		// Numbers must be prefixed with 0 to match the length of the field
		// so I guess floats should also be padded
		s := strings.TrimLeft(string(value), "0")

		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		return f, nil
	})

	MustRegister(FieldTypeDate, func(value []byte, desc *FieldDesc) (interface{}, error) {
		if err := verifyLength(desc, value); err != nil {
			return nil, err
		}

		// unfortunately GDT is one of those specifications that doesn't give a shit
		// on timezone support. We use time.Local here for parsing as a best practice.
		// Users that need a different timezone must re-implement this function or
		// add a dedicated type (and use a different TypeRegistry).

		parse := func(b []byte) (int, error) {
			i, err := strconv.ParseInt(string(b), 10, 0)
			return int(i), err
		}

		day, err := parse(value[:2])
		if err != nil {
			return nil, err
		}
		month, err := parse(value[2:4])
		if err != nil {
			return nil, err
		}
		year, err := parse(value[4:])
		if err != nil {
			return nil, err
		}

		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil
	})

	MustRegister(FieldTypeTime, func(value []byte, desc *FieldDesc) (interface{}, error) {
		if err := verifyLength(desc, value); err != nil {
			return nil, err
		}

		// same issue as in FieldTypeDate

		parse := func(b []byte, mul time.Duration) (time.Duration, error) {
			i, err := strconv.ParseInt(string(b), 10, 0)
			return time.Duration(i) * mul, err
		}

		hour, err := parse(value[:2], time.Hour)
		if err != nil {
			return time.Duration(0), err
		}
		min, err := parse(value[2:4], time.Minute)
		if err != nil {
			return time.Duration(0), err
		}
		sec, err := parse(value[4:6], time.Second)
		if err != nil {
			return time.Duration(0), err
		}
		return (hour + min + sec), nil
	})
}
