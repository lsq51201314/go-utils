package gzlib

import (
	"bytes"
	"compress/zlib"
	"io"

	"github.com/lsq51201314/go-utils/gbytes"
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
	return in.Bytes()[2:] //去除78da
}

// 解压数据
func UnCompress(src []byte) []byte {
	s := gbytes.BytesCombine([]byte{0x78, 0xda}, src) //加上78da
	b := bytes.NewReader(s)
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil
	}
	var out bytes.Buffer
	if _, err := io.Copy(&out, r); err != nil {
		return nil
	}
	if err := r.Close(); err != nil {
		return nil
	}
	return out.Bytes()
}
