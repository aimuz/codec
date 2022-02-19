package json

import (
	"encoding/json"

	"github.com/aimuz/codec"
)

// Codec that encodes to and decodes from JSON.
var Codec codec.Codec = new(c)

type c struct{}

// Marshal returns the wire format of v.
func (c) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal parses the wire format into v.
func (c) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

// Name return codec name
func (c) Name() string {
	return "json"
}
