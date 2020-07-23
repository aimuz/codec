package snappy

import (
	"github.com/aimuz/codec"
	"github.com/golang/snappy"
)

// Codec that encodes to and decodes from SNAPPY.
type snappyCodec struct {
	codec codec.Codec
}

var _ codec.Codec = new(snappyCodec)

// NewSnappyCodec ...
func NewSnappyCodec(codec codec.Codec) *snappyCodec {
	return &snappyCodec{
		codec: codec,
	}
}

func (m snappyCodec) Marshal(v interface{}) ([]byte, error) {
	b, err := m.codec.Marshal(v)
	if err != nil {
		return nil, err
	}
	return snappy.Encode(nil, b), nil
}

func (m snappyCodec) Unmarshal(b []byte, v interface{}) error {
	var err error
	b, err = snappy.Decode(nil, b)
	if err != nil {
		return err
	}
	return m.codec.Unmarshal(b, v)
}

func (m snappyCodec) Name() string {
	return "snappy+" + m.codec.Name()
}
