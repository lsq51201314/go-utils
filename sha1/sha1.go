package sha1

import (
	"crypto/sha1"
	"encoding/hex"
)

// 数据摘要
func Get(data []byte) string {
	buf := sha1.Sum(data)
	return hex.EncodeToString(buf[:])
}
