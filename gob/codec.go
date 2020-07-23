package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/aimuz/codec"
)

const name = "json"

// Codec that encodes to and decodes from GOB.
var Codec = new(gobCodec)

type gobCodec int

var _ codec.Codec = new(gobCodec)

// Marshal returns the wire format of v.
func (g gobCodec) Marshal(v interface{}) ([]byte, error) {
	var w bytes.Buffer
	err := gob.NewEncoder(&w).Encode(v)
	return w.Bytes(), err
}

// Unmarshal parses the wire format into v.
func (g gobCodec) Unmarshal(b []byte, v interface{}) error {
	return gob.NewDecoder(bytes.NewReader(b)).Decode(v)
}

// Name return codec name
func (g gobCodec) Name() string {
	return name
}
