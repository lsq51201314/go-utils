package utils

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliyunOSSOptions struct {
	AccessKeyId     string
	AccessKeySecret string
	OssEndpoint     string
	OssBucketName   string
}

type AliyunOSS struct {
	client *oss.Client
	bucket *oss.Bucket
}

func NewAliyunOSS(options AliyunOSSOptions) (aliyun AliyunOSS, err error) {
	if aliyun.client, err = oss.New(
		options.OssEndpoint,
		options.AccessKeyId,
		options.AccessKeySecret); err != nil {
		return
	}
	if aliyun.bucket, err = aliyun.client.Bucket(options.OssBucketName); err != nil {
		return
	}
	return
}

func (a *AliyunOSS) getName(s []byte) (name string, err error) {
	h := sha1.New()
	if _, err = h.Write(s); err != nil {
		return
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func (a *AliyunOSS) Upload(path string, data []byte, name ...string) (res string, err error) {
	if path[:1] != "/" {
		path = "/" + path
	}
	//获取名称
	var file string
	if len(name) == 0 {
		if file, err = a.getName(data); err != nil {
			return
		}
	} else {
		file = strings.TrimSpace(name[0])
	}
	dir := path[1:] + "/" + file
	//使用修改的方式是为了防止自定义名称无法覆盖原数据
	err = a.bucket.PutObject(dir, bytes.NewReader(data))
	fmt.Println(dir)
	return
}

func (a *AliyunOSS) Download(path string) (data []byte, err error) {
	if path[:1] != "/" {
		path = "/" + path
	}

	var body io.ReadCloser
	if body, err = a.bucket.GetObject(path[1:]); err != nil {
		return
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(body)
	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, body); err != nil {
		return
	}
	return buf.Bytes(), nil
}

func (a *AliyunOSS) Delete(path string) (err error) {
	if path[:1] != "/" {
		path = "/" + path
	}

	err = a.bucket.DeleteObject(path[1:])
	return
}
