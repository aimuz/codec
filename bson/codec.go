package bson

import (
	"github.com/aimuz/codec"
	"go.mongodb.org/mongo-driver/bson"
)

const name = "bson"

// Codec that encodes to and decodes from BSON.
var Codec = new(bsonCodec)

type bsonCodec int

var _ codec.Codec = new(bsonCodec)

// Marshal returns the wire format of v.
func (c bsonCodec) Marshal(v interface{}) ([]byte, error) {
	return bson.Marshal(v)
}

// Unmarshal parses the wire format into v.
func (c bsonCodec) Unmarshal(b []byte, v interface{}) error {
	return bson.Unmarshal(b, v)
}

// Name return codec name
func (c bsonCodec) Name() string {
	return name
}
