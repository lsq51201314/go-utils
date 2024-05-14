package utils

import (
	"bytes"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 文件信息
type YunFile struct {
	Path         string
	Size         int64
	LastModified string
}

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

// 上传文件
func (a *YunOSS) Upload(path string, data []byte) (err error) {
	if path[:1] != "/" {
		path = "/" + path
	}
	//存储文件
	err = a.bucket.PutObject(path[1:], bytes.NewReader(data))
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
func (a *YunOSS) Query(prefix string, size int, next ...string) (files []YunFile, token string, err error) {
	if size <= 0 {
		size = 100
	}
	if size > 1000 {
		size = 1000
	}
	if prefix[:1] != "/" {
		prefix = "/" + prefix
	}
	files = []YunFile{}
	token = ""
	if len(next) > 0 {
		token = next[0]
	}
	var lsRes oss.ListObjectsResultV2
	a.bucket.ListObjects()
	if lsRes, err = a.bucket.ListObjectsV2(
		oss.ContinuationToken(token),
		oss.MaxKeys(size),
		oss.Prefix(prefix[1:])); err != nil {
		return
	}
	for _, object := range lsRes.Objects {
		if object.Key != prefix[1:] {
			files = append(files, YunFile{
				Path:         object.Key,
				Size:         object.Size,
				LastModified: object.LastModified.In(location).Format("2006-01-02 15:04:05"),
			})
		}
	}
	if lsRes.IsTruncated {
		token = lsRes.NextContinuationToken
	} else {
		token = ""
	}
	return
}
