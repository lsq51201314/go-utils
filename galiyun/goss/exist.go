package goss

import "context"

// 文件存在
func (t *OSS) Exist(fileName string) (bool, error) {
	return t.ossClient.IsObjectExist(context.TODO(), t.bucketName, fileName)
}
