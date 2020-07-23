package json

import (
	"encoding/json"
	"github.com/aimuz/codec"
)

const name = "json"

// Codec that encodes to and decodes from JSON.
var Codec = new(jsonCodec)

type jsonCodec int

var _ codec.Codec = new(jsonCodec)

func (j jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j jsonCodec) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

func (j jsonCodec) Name() string {
	return name
}
