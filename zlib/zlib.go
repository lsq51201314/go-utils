package zlib

import (
	"bytes"
	"compress/zlib"
	"io"
)

// 压缩数据
func Compress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	if _, err := w.Write(src); err != nil {
		return nil
	}
	if err := w.Close(); err != nil {
		return nil
	}
	return in.Bytes()
}

// 解压数据
func UnCompress(src []byte) []byte {
	b := bytes.NewReader(src)
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil
	}
	var out bytes.Buffer
	if _, err := io.Copy(&out, r); err != nil {
		return nil
	}
	return out.Bytes()
}
