package utils

import (
	"bytes"
	"compress/zlib"
	"io"
)

type Zlib struct{}

func (z Zlib) Compress(src []byte) (data []byte, err error) {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	if _, err = w.Write(src); err != nil {
		return
	}
	if err = w.Close(); err != nil {
		return
	}
	data = in.Bytes()
	return
}

func (z Zlib) UnCompress(src []byte) (data []byte, err error) {
	b := bytes.NewReader(src)
	var out bytes.Buffer
	var r io.ReadCloser
	if r, err = zlib.NewReader(b); err != nil {
		return
	}
	if _, err = io.Copy(&out, r); err != nil {
		return
	}
	data = out.Bytes()
	return
}
