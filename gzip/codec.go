package gzip

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	"github.com/aimuz/codec"
)

type c struct {
	level int
	codec codec.Codec
}

var _ codec.Codec = (*c)(nil)

// NewGzipCodecWith ...
func NewGzipCodecWith(coder codec.Codec) codec.Codec {
	return &c{
		codec: coder,
	}
}

// NewGzipCodecLevelWith ...
func NewGzipCodecLevelWith(coder codec.Codec, level int) codec.Codec {
	return &c{
		codec: coder,
		level: level,
	}
}

// Marshal returns the wire format of v.
func (c c) Marshal(v interface{}) ([]byte, error) {
	b, err := c.codec.Marshal(v)
	if err != nil {
		return nil, err
	}
	var w bytes.Buffer
	var gw *gzip.Writer
	if c.level > 0 {
		gw, err = gzip.NewWriterLevel(&w, c.level)
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
func (c c) Unmarshal(b []byte, v interface{}) error {
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
	return c.codec.Unmarshal(b, v)
}

// Name return codec name
func (c c) Name() string {
	return "gzip+" + c.codec.Name()
}
