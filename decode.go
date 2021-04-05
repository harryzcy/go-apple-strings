package applestrings

import (
	"bufio"
	"errors"
	"io"
	"reflect"
	"strings"
)

var (
	ErrInvalidSyntax = errors.New("invalid syntax")
)

type Decoder struct {
	reader *bufio.Reader
}

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

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{reader: bufio.NewReader(r)}
}

// getString extract the string from the quote or unquote form, allowing leading or trailing spaces.
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
