package gdt

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

// DecoderOption can be passed to NewParser to configure different aspects
// of the GDT Parser.
type DecoderOption func(p *Decoder) error

// WithDefaultChartset configures the default character set used by the
// decoder. Without this option, the decoder is configured to use
// IBM-8bit encoding (Codepage 437)A.
func WithDefaultCharset(m *charmap.Charmap) DecoderOption {
	return func(d *Decoder) error {
		d.charDecoder = m.NewDecoder()
		return nil
	}
}

// Decoder is capable of decoding GDT data.
type Decoder struct {
	r           io.Reader
	charDecoder *encoding.Decoder
}

type FieldID uint

func (f FieldID) String() string { return fmt.Sprintf("%04d", f) }

type Line struct {
	FieldID FieldID
	Content []byte
}

func NewDecoder(r io.Reader, options ...DecoderOption) (*Decoder, error) {
	d := &Decoder{
		charDecoder: charmap.CodePage437.NewDecoder(),
	}

	for _, opt := range options {
		if err := opt(d); err != nil {
			return nil, err
		}
	}

	return d, nil
}

type File struct {
	Charset encoding.Encoding

	// Lines holds all lines of the file.
	Lines []Line
}

// NewFile creates a File from lines by parsing important ones like
// 9206 which specify the character set to use.
func NewFile(lines Lines) *File {
	var cm encoding.Encoding = charmap.CodePage437
	if f, ok := lines.GetField(9206); ok {
		switch f.Content[0] {
		case '1':
			cm = encoding.Nop // TODO(ppacher): ASCII-7-bit encoding
		case '2':
			cm = charmap.CodePage437
		case '3':
			cm = charmap.Windows1252
		}
	}
	return &File{
		Charset: cm,
		Lines:   lines,
	}
}

// Lines is a slice of lines and provides utility methods.
type Lines []Line

// GetField returns the field with Field-ID id.
func (lines Lines) GetField(id FieldID) (Line, bool) {
	for _, l := range lines {
		if l.FieldID == id {
			return l, true
		}
	}
	return Line{}, false
}

// ReadLines reads all GDT lines stored in the underlying
// reader. Note that the line data returned is not decoded.
func (dec *Decoder) ReadLines() (File, error) {
	var content Lines
	for {
		l, err := dec.NextLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
			}
			return *NewFile(content), err
		}
		content = append(content, *l)
	}
}

// NextLine reads the next line from the underlying reader.
func (dec *Decoder) NextLine() (*Line, error) {
	// read the first 3 bytes which contain the line length as
	// as a string.
	var lenbuf [3]byte
	if _, err := io.ReadAtLeast(dec.r, lenbuf[:], 3); err != nil {
		return nil, err
	}

	// decode the line length
	l, err := dec.decodeInt(lenbuf[:])
	if err != nil {
		return nil, fmt.Errorf("failed to decode line length: %w", err)
	}

	// read the rest of the line
	var buf = make([]byte, l)
	if _, err := io.ReadFull(dec.r, buf); err != nil {
		return nil, err
	}

	// the last two bytes of the line MUST be CRLF
	if buf[l-2] != '\r' || buf[l-1] != '\n' {
		return nil, fmt.Errorf("unexpected line endings 0x%02x 0x%02x", buf[l-2], buf[l-1])
	}

	fieldID, err := dec.decodeInt(buf[:4])
	if err != nil {
		return nil, fmt.Errorf("failed to decode field identifier: %w", err)
	}

	content := buf[4 : len(buf)-2]
	return &Line{
		FieldID: FieldID(fieldID),
		Content: content,
	}, nil
}

func (dec *Decoder) decodeInt(in []byte) (int, error) {
	decoded, err := dec.charDecoder.Bytes(in)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(string(decoded), 10, 0)
	return int(i), err
}
