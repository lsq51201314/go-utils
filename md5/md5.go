package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// 数据摘要
func Get(data []byte) string {
	buf := md5.Sum(data)
	return hex.EncodeToString(buf[:])
}
