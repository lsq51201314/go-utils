package utils

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 存储实例
type YunOSS struct {
	client *oss.Client
	bucket *oss.Bucket
}

// 新建实例
func NewYunOSS(accessKeyId, accessKeySecret, ossBucketName string, ossEndpoint ...string) (yun YunOSS, err error) {
	endpoint := "oss-cn-shenzhen.aliyuncs.com"
	if len(ossEndpoint) > 0 {
		endpoint = ossEndpoint[0]
	}
	if yun.client, err = oss.New(endpoint, accessKeyId, accessKeySecret); err != nil {
		return
	}
	if yun.bucket, err = yun.client.Bucket(ossBucketName); err != nil {
		return
	}
	return
}

func (a *YunOSS) getName(s []byte) (name string, err error) {
	h := sha1.New()
	if _, err = h.Write(s); err != nil {
		return
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// 上传文件
func (a *YunOSS) Upload(path string, data []byte, name ...string) (res string, err error) {
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
	//存储文件
	err = a.bucket.PutObject(dir, bytes.NewReader(data))
	return
}

// 下载文件
func (a *YunOSS) Download(path string) (data []byte, err error) {
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

// 删除文件
func (a *YunOSS) Delete(path string) (err error) {
	if path[:1] != "/" {
		path = "/" + path
	}

	err = a.bucket.DeleteObject(path[1:])
	return
}

// 枚举文件
func (a *YunOSS) Query(prefix string, size int, next ...string) (files []string, token string, err error) {
	if size <= 0 {
		size = 100
	}
	if size > 1000 {
		size = 1000
	}
	if prefix[:1] != "/" {
		prefix = "/" + prefix
	}
	files = []string{}
	token = ""
	if len(next) > 0 {
		token = next[0]
	}
	var lsRes oss.ListObjectsResultV2
	a.bucket.ListObjects()
	if lsRes, err = a.bucket.ListObjectsV2(oss.ContinuationToken(token), oss.MaxKeys(size), oss.Prefix(prefix[1:])); err != nil {
		return
	}
	for _, object := range lsRes.Objects {
		if object.Key != prefix[1:] {
			files = append(files, object.Key)
		}
	}
	if lsRes.IsTruncated {
		token = lsRes.NextContinuationToken
	} else {
		token = ""
	}
	return
}
