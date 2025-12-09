package goss

import (
	"context"
	"io"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
)

// 获取文件
func (t *Goss) Get(fileName string) ([]byte, error) {
	request := &oss.GetObjectRequest{
		Bucket: oss.Ptr(t.bucketName),
		Key:    oss.Ptr(fileName),
	}

	result, err := t.ossClient.GetObject(context.TODO(), request)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	return io.ReadAll(result.Body)
}
