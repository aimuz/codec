package xml

import (
	"encoding/xml"
	"github.com/aimuz/codec"
)

const name = "xml"

// Codec that encodes to and decodes from XML.
var Codec = new(xmlCodec)

type xmlCodec int

var _ codec.Codec = new(xmlCodec)

// Marshal returns the wire format of v.
func (j xmlCodec) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

// Unmarshal parses the wire format into v.
func (j xmlCodec) Unmarshal(b []byte, v interface{}) error {
	return xml.Unmarshal(b, v)
}

// Name return codec name
func (j xmlCodec) Name() string {
	return name
}

