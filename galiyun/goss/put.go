package goss

import (
	"bytes"
	"context"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/lsq51201314/go-utils/gsha"
)

// 上传文件
func (t *OSS) Put(data []byte) (string, error) {
	fileName := gsha.Sha1Str(data)
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
