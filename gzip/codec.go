package gzip

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	"github.com/aimuz/codec"
)

// Codec that encodes to and decodes from GZIP.
type gzipCodec struct {
	level int
	codec codec.Codec
}

var _ codec.Codec = new(gzipCodec)

// NewGzipCodecWith ...
func NewGzipCodecWith(coder codec.Codec) codec.Codec {
	return &gzipCodec{
		codec: coder,
	}
}

// NewGzipCodecLevelWith ...
func NewGzipCodecLevelWith(coder codec.Codec, level int) codec.Codec {
	return &gzipCodec{
		codec: coder,
		level: level,
	}
}

// Marshal returns the wire format of v.
func (m gzipCodec) Marshal(v interface{}) ([]byte, error) {
	b, err := m.codec.Marshal(v)
	if err != nil {
		return nil, err
	}
	var w bytes.Buffer
	var gw *gzip.Writer
	if m.level > 0 {
		gw, err = gzip.NewWriterLevel(&w, m.level)
		if err != nil {
			return nil, err
		}
	} else {
		gw = gzip.NewWriter(&w)
	}
	_, err = gw.Write(b)
	if err != nil {
		return nil, err
	}
	err = gw.Flush()
	return w.Bytes(), err
}

// Unmarshal parses the wire format into v.
func (m gzipCodec) Unmarshal(b []byte, v interface{}) error {
	var (
		err error
		r   *gzip.Reader
	)
	r, err = gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer r.Close()
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return m.codec.Unmarshal(b, v)
}

// Name return codec name
func (m gzipCodec) Name() string {
	return "gzip+" + m.codec.Name()
}
