package gaes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const defaultKey string = "No.9"

// 加密数据
func Enc(data []byte, passwd ...string) []byte {
	key := defaultKey
	if len(passwd) > 0 {
		key = passwd[0]
	}
	key = getMD5Key(key)
	iv := getIV(key)
	key = string(zerosPadding([]byte(key), aes.BlockSize))
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil
	}
	src := pKCS7Padding(data, block.BlockSize())
	encryptData := make([]byte, len(src))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(encryptData, src)
	return encryptData
}

// 解密数据
func Dec(data []byte, passwd ...string) []byte {
	key := defaultKey
	if len(passwd) > 0 {
		key = passwd[0]
	}
	key = getMD5Key(key)
	iv := getIV(key)
	key = string(zerosPadding([]byte(key), aes.BlockSize))
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil
	}
	decryptData := make([]byte, len(data))
	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(decryptData, data)
	original, err := pKCS7UnPadding(decryptData)
	if err != nil {
		return nil
	}
	return original
}

func pKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func pKCS7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return src, fmt.Errorf("数据为空")
	}
	unpadding := int(src[length-1])
	if length < unpadding {
		return src, fmt.Errorf("长度错误")
	}
	return src[:(length - unpadding)], nil
}

func zerosPadding(src []byte, blockSize int) []byte {
	rem := len(src) % blockSize
	if rem == 0 {
		return src
	}
	return append(src, bytes.Repeat([]byte{0}, blockSize-rem)...)
}

func getBlockSize(key string) string {
	if len(key) > aes.BlockSize {
		return key[:aes.BlockSize]
	}
	return key
}

func getMD5Key(key string) string {
	passwd := md5.Sum([]byte(key))
	return hex.EncodeToString(passwd[:])
}

func getIV(key string) string {
	hashedKey := sha256.Sum256([]byte(key))
	return getBlockSize(hex.EncodeToString(hashedKey[:]))
}
