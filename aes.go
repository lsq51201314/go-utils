package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
)

// 加密实例
type Aes struct {
	key [16]byte
}

// 新建实例
func NewAes(passwd string) *Aes {
	return &Aes{key: md5.Sum([]byte(passwd))}
}

// 加密数据
func (a *Aes) Enc(src []byte) (data []byte, err error) {
	var block cipher.Block
	if block, err = aes.NewCipher(a.key[:]); err != nil {
		return
	}
	blockSize := block.BlockSize()
	src = pkcs5Padding(src, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, a.key[:blockSize])
	data = make([]byte, len(src))
	blockMode.CryptBlocks(data, src)
	return
}

// 解密数据
func (a *Aes) Dec(src []byte) (data []byte, err error) {
	var block cipher.Block
	if block, err = aes.NewCipher(a.key[:]); err != nil {
		return
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, a.key[:blockSize])
	data = make([]byte, len(src))
	blockMode.CryptBlocks(data, src)
	data = pkcs5UnPadding(data)
	return
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
