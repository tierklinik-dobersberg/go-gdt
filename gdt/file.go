package gdt

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

type File struct {
	charset      encoding.Encoding
	typeRegistry *TypeRegistry

	// Lines holds all lines of the file.
	Lines []Line
}

// NewFile creates a File from lines by parsing important ones like
// 9206 which specify the character set to use.
func NewFile(lines Lines) *File {
	file := NewFileWithRegistry(lines, DefaultRegistry)
	if val, err := file.IntField(FieldCharacterSet); err == nil { // we drop the error otherwise and let the user handle that
		switch val {
		case '1':
			file.charset = encoding.Nop // TODO(ppacher): ASCII-7-bit encoding
		case '2':
			file.charset = charmap.CodePage437
		case '3':
			file.charset = charmap.Windows1252
		}
	}

	return file
}

// NewFileWithRegistry creates a new GDT file representation that uses
// reg for known types. Users most commonly only need to use NewFile
// instead of NewFileWithRegistry.
func NewFileWithRegistry(lines Lines, reg *TypeRegistry) *File {
	if reg == nil {
		reg = DefaultRegistry
	}
	return &File{
		charset:      charmap.CodePage437,
		typeRegistry: reg,
		Lines:        lines,
	}
}

// Lines is a slice of lines and provides utility methods.
type Lines []Line

// Field returns the decoded value of the field record described
// by desc. The first line that matches is decoded and returned.
// Use Fields() to get a slice of all lines that match.
func (file *File) Field(desc FieldDesc) (interface{}, error) {
	for _, l := range file.Lines {
		if l.FieldID == desc.ID {
			// decode the content using the configured
			// charset
			content, err := file.charset.NewDecoder().Bytes(l.Content)
			if err != nil {
				return nil, err
			}
			// use the type registry to actually decode the field
			return file.typeRegistry.ParseValue(content, &desc)
		}
	}
	return nil, ErrFieldNotFound
}

// IntField is like Field() but casts the returned interface to
// int.
func (file *File) IntField(desc FieldDesc) (int, error) {
	val, err := file.Field(desc)
	if err != nil {
		return 0, err
	}
	i, ok := val.(int)
	if !ok {
		return 0, ErrInvalidType
	}
	return i, nil
}

// FloatField is like Field() but casts the returned interface to
// float64.
func (file *File) FloatField(desc FieldDesc) (float64, error) {
	val, err := file.Field(desc)
	if err != nil {
		return 0, err
	}
	i, ok := val.(float64)
	if !ok {
		return 0, ErrInvalidType
	}
	return i, nil
}

// StringField is like Field() but casts the returned interface to
// string.
func (file *File) StringField(desc FieldDesc) (string, error) {
	val, err := file.Field(desc)
	if err != nil {
		return "", err
	}
	i, ok := val.(string)
	if !ok {
		return "", ErrInvalidType
	}
	return i, nil
}

// Fields is like Field but returns a slice containing the value
// for each occurance of desc in file.
func (file *File) Fields(desc FieldDesc) ([]interface{}, error) {
	var result []interface{}
	for _, l := range file.Lines {
		if l.FieldID == desc.ID {
			content, err := file.charset.NewDecoder().Bytes(l.Content)
			if err != nil {
				return result, err
			}
			val, err := file.typeRegistry.ParseValue(content, &desc)
			if err != nil {
				return result, err
			}
			result = append(result, val)
		}
	}
	return result, nil
}
