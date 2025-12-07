package goss

import (
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

// 对象存储
type OSS struct {
	endpoint        string
	region          string
	bucketName      string
	accessKeyId     string
	accessKeySecret string
	ossClient       *oss.Client
}

// 新建
func New(endpoint, region, bucketName, accessKeyId, accessKeySecret string) *OSS {
	obj := OSS{
		endpoint:        endpoint,
		region:          region,
		bucketName:      bucketName,
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
		ossClient: oss.NewClient(oss.LoadDefaultConfig().
			WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				accessKeyId,
				accessKeySecret)).
			WithRegion(region).
			WithEndpoint(endpoint)),
	}
	return &obj
}
