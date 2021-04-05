package applestrings

import (
	"bufio"
	"errors"
	"io"
	"reflect"
	"strings"
)

var (
	// ErrInvalidSyntax is returned and the syntax of input data is invalid.
	ErrInvalidSyntax = errors.New("invalid syntax")
)

// Decoder reads the decoder stream.
type Decoder struct {
	reader *bufio.Reader
}

// Decode reads the decoder stream to find key-value pairs,
// similar to how Unmarshal works.
// The implementation complies with Apple's documentation at
// https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/LoadingResources/Strings/Strings.html
func (d *Decoder) Decode(v interface{}) error {
	for {
		line, err := d.reader.ReadString(byte('\n'))
		if err != nil && err != io.EOF {
			return err
		}

		line = strings.Trim(line, " \n")
		if line == "" {
			if err == io.EOF {
				break
			}
			continue
		}

		// skip comment
		if strings.HasPrefix(line, "/*") {
			if strings.HasSuffix(line, "*/") {
				continue
			} else {
				return ErrInvalidSyntax
			}
		}

		// parse line of the format "x" = "y";
		if !strings.HasSuffix(line, ";") {
			return ErrInvalidSyntax
		}
		line = line[:len(line)-1] // remove trailing semicolon

		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			return ErrInvalidSyntax
		}

		key := getString(parts[0])
		value := getString(parts[1])
		d.set(v, key, value)

		if err == io.EOF {
			break
		}
	}

	return nil
}

// set finds the struct field by key and sets it with the value.
func (d *Decoder) set(v interface{}, key, value string) {
	p := reflect.ValueOf(v)
	if p.Kind() != reflect.Ptr {
		return
	}

	t := reflect.Indirect(p)
	if t.Kind() != reflect.Struct {
		return
	}

	f := t.FieldByName(key)
	if !f.CanSet() || f.Kind() != reflect.String {
		return
	}

	f.SetString(value)
}

// NewDecoder returns a Decoder that reads Apple strings file.
// NewDecoder requires a stream reader, r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{reader: bufio.NewReader(r)}
}

// getString extract the string from the quote or unquote form,
// allowing leading or trailing spaces.
func getString(s string) string {
	str := strings.Trim(s, " ")

	if str[0] == '"' {
		str = str[1:]
	}
	if str[len(str)-1] == '"' {
		str = str[:len(str)-1]
	}

	return strings.ReplaceAll(str, `\"`, `"`)
}
