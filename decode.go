package applestrings

import "io"

type Decoder struct {
	reader io.ReadSeeker
}

func (d *Decoder) Decode(v interface{}) error {
	return nil
}

func NewDecoder(r io.ReadSeeker) *Decoder {
	return &Decoder{reader: r}
}
