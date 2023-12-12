package utils

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
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

func (a *AliyunOSS) Upload(data []byte, path ...string) (name string, err error) {
	//获取名称
	if len(path) == 0 {
		if name, err = a.getName(data); err != nil {
			return
		}
	} else {
		name = strings.TrimSpace(path[0])
	}
	//使用修改的方式是为了防止自定义名称无法覆盖原数据
	err = a.bucket.PutObject(name, bytes.NewReader(data))
	return
}

func (a *AliyunOSS) Download(name string) (data []byte, err error) {
	var body io.ReadCloser
	if body, err = a.bucket.GetObject(name); err != nil {
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

func (a *AliyunOSS) Delete(name string) (err error) {
	err = a.bucket.DeleteObject(name)
	return
}
