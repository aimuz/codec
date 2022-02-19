package snappy

import (
	"github.com/aimuz/codec"
	"github.com/golang/snappy"
)

// Codec that encodes to and decodes from SNAPPY.
type c struct {
	codec codec.Codec
}

var _ codec.Codec = (*c)(nil)

// NewSnappyCodecWith ...
func NewSnappyCodecWith(codec codec.Codec) codec.Codec {
	return &c{
		codec: codec,
	}
}

// Marshal returns the wire format of v.
func (c c) Marshal(v interface{}) ([]byte, error) {
	b, err := c.codec.Marshal(v)
	if err != nil {
		return nil, err
	}
	return snappy.Encode(nil, b), nil
}

// Unmarshal parses the wire format into v.
func (c c) Unmarshal(b []byte, v interface{}) error {
	var err error
	b, err = snappy.Decode(nil, b)
	if err != nil {
		return err
	}
	return c.codec.Unmarshal(b, v)
}

// Name return codec name
func (c c) Name() string {
	return "snappy+" + c.codec.Name()
}
