package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/aimuz/codec"
)

const name = "json"

// Codec that encodes to and decodes from GOB.
var Codec codec.Codec = new(c)

type c struct{}

// Marshal returns the wire format of v.
func (c) Marshal(v interface{}) ([]byte, error) {
	var w bytes.Buffer
	err := gob.NewEncoder(&w).Encode(v)
	return w.Bytes(), err
}

// Unmarshal parses the wire format into v.
func (c) Unmarshal(b []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewReader(b)).Decode(v)
}

// Name return codec name
func (c) Name() string {
	return name
}
