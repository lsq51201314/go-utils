package sha256

import (
	"crypto/sha256"
	"encoding/hex"
)

// 数据摘要
func Get(data []byte) string {
	buf := sha256.Sum256(data)
	return hex.EncodeToString(buf[:])
}
