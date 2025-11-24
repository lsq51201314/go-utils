package oss

import (
	"bytes"
	"context"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/lsq51201314/go-utils/sha1"
)

// 上传文件
func (t *OSS) Put(data []byte) (string, error) {
	fileName := sha1.Get(data)
	request := &oss.PutObjectRequest{
		Bucket: oss.Ptr(t.bucketName),
		Key:    oss.Ptr(fileName),
		Body:   bytes.NewBuffer(data),
	}
	if _, err := t.ossClient.PutObject(context.TODO(), request); err != nil {
		return "", err
	}
	return fileName, nil
}
