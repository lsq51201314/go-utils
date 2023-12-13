package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"os"
	"strings"
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

func (y *YunFile) Upload(path string, data []byte, name ...string) (res string, err error) {
	if path[:1] != "/" {
		path = "/" + path
	}
	//创建目录
	dir := y.dir + path
	var exist bool
	if exist, err = y.exist(dir); err != nil {
		return
	}
	if !exist {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return
		}
	}
	dir += "/"
	//获取名称
	var file string
	if len(name) == 0 {
		if file, err = y.getName(data); err != nil {
			return
		}
	} else {
		file = strings.TrimSpace(name[0])
	}
	//存储文件
	err = os.WriteFile(dir+file, data, os.ModePerm)
	res = path + "/" + file
	return
}

func (y *YunFile) Download(path string) (data []byte, err error) {
	if path[:1] != "/" {
		path = "/" + path
	}
	//判断存在
	var exist bool
	if exist, err = y.exist(y.dir + path); err != nil {
		return
	}
	if !exist {
		err = errors.New("file not found")
		return
	}
	//读取文件
	data, err = os.ReadFile(y.dir + path)
	return
}

func (y *YunFile) Delete(path string) (err error) {
	if path[:1] != "/" {
		path = "/" + path
	}
	//判断存在
	var exist bool
	if exist, err = y.exist(y.dir + path); err != nil {
		return
	}
	if !exist {
		err = errors.New("file not found")
		return
	}
	//删除文件
	err = os.Remove(y.dir + path)
	return
}
