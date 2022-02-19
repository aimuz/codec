package protobuf

import (
	"github.com/aimuz/codec"
	"google.golang.org/protobuf/proto"
)

// Codec that encodes to and decodes from Protobuf.
var Codec codec.Codec = new(c)

type c struct{}

func (c) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (c) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

func (c) Name() string {
	return "protobuf"
}
