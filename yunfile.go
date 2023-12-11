package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"os"
)

type YunFile struct {
	dir string
}

func NewYunFile() (y YunFile, err error) {
	y.dir = "./assets"
	var exist bool
	if exist, err = y.exist(y.dir); err != nil {
		return
	}
	if !exist {
		err = os.Mkdir(y.dir, os.ModePerm)
	}
	return
}

func (y *YunFile) exist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (y *YunFile) getName(s []byte) (name string, err error) {
	h := sha1.New()
	if _, err = h.Write(s); err != nil {
		return
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func (y *YunFile) Upload(data []byte) (name string, err error) {
	//获取名称
	if name, err = y.getName(data); err != nil {
		return
	}
	//判断存在
	var exist bool
	if exist, err = y.exist(y.dir + "/" + name); err != nil {
		return
	}
	//存储文件
	if !exist {
		err = os.WriteFile(y.dir+"/"+name, data, os.ModePerm)
	}
	return
}

func (y *YunFile) Download(name string) (data []byte, err error) {
	//判断存在
	var exist bool
	if exist, err = y.exist(y.dir + "/" + name); err != nil {
		return
	}
	if !exist {
		err = errors.New("file not found")
		return
	}
	//读取文件
	data, err = os.ReadFile(y.dir + "/" + name)
	return
}

func (y *YunFile) Delete(name string) (err error) {
	err = os.Remove(y.dir + "/" + name)
	return
}
