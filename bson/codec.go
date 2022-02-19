package bson

import (
	"github.com/aimuz/codec"
	"go.mongodb.org/mongo-driver/bson"
)

// Codec that encodes to and decodes from BSON.
var Codec codec.Codec = new(c)

type c struct{}

// Marshal returns the wire format of v.
func (c) Marshal(v interface{}) ([]byte, error) {
	return bson.Marshal(v)
}

// Unmarshal parses the wire format into v.
func (c) Unmarshal(b []byte, v interface{}) error {
	return bson.Unmarshal(b, v)
}

// Name return codec name
func (c) Name() string {
	return "bson"
}
