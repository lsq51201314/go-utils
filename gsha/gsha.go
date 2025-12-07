package gsha

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

//数据摘要
func Sha1(data []byte) []byte {
	buf := sha1.Sum(data)
	return buf[:]
}

//数据摘要
func Sha256(data []byte) []byte {
	buf := sha256.Sum256(data)
	return buf[:]
}

//数据摘要
func Sha1Str(data []byte) string {
	buf := sha1.Sum(data)
	return hex.EncodeToString(buf[:])
}

//数据摘要
func Sha256Str(data []byte) string {
	buf := sha256.Sum256(data)
	return hex.EncodeToString(buf[:])
}
