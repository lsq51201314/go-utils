package zlib

import (
	"bytes"
	"compress/zlib"
	"io"
)

// 压缩数据
func Compress(src []byte) []byte {
	var in bytes.Buffer
	w, err := zlib.NewWriterLevel(&in, zlib.BestCompression) //压缩等级9 最佳压缩
	if err != nil {
		return nil
	}
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
	defer r.Close()

	var out bytes.Buffer
	if _, err := io.Copy(&out, r); err != nil {
		return nil
	}
	return out.Bytes()
}
