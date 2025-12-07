package gfile

import (
	"os"
	"path/filepath"
)

// 获取列表
func GetFiles(dir string) []string {
	if dir[len(dir)-1:] != "/" {
		dir += "/"
	}
	fd, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	files := make([]string, 0)
	for _, fi := range fd {
		fullname := dir + fi.Name()
		if fi.IsDir() {
			var temp []string
			if temp = GetFiles(fullname + "/"); temp == nil {
				return nil
			}
			files = append(files, temp...)
		} else {
			files = append(files, fullname)
		}
	}
	return files
}

// 创建多级目录
func CreateMutiDir(fullPath string, isFile ...bool) error {
	var path string
	if len(isFile) > 0 && isFile[0] {
		path = filepath.Dir(fullPath)
	} else {
		path = fullPath
	}
	if !IsExist(path) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// 是否存在
func IsExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return os.IsExist(err)
	}
	return true
}
