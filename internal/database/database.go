package database

import (
	"log"
	"time"

	. "vault-server/cmd/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var client *s3.S3
var bucket *string

func NewS3Client() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(Cfg.Region),
		Endpoint:    aws.String(Cfg.Endpoint),
		Credentials: credentials.NewStaticCredentials(Cfg.AccessKey, Cfg.SecretKey, ""),
	})
	if err != nil {
		log.Fatal(err)
	}
	bucket = aws.String(Cfg.BucketName)
	client = s3.New(sess)
}

func ListObjects(prefix string) (*s3.ListObjectsV2Output, error) {
	res, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: bucket,
		Prefix: aws.String(prefix),
	})
	return res, err
}

func ObjectGetUrl(key string) (string, error) {
	res, _ := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket:                     bucket,
		Key:                        aws.String(key),
		ResponseContentDisposition: aws.String("attachment"),
	})
	url, err := res.Presign(60 * time.Minute)
	return url, err
}

func ObjectPutUrl(key string) (string, error) {
	res, _ := client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: bucket,
		Key:    aws.String(key),
	})
	url, err := res.Presign(3 * time.Minute)
	return url, err
}

func DeleteObject(key string) error {
	_, err := client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: bucket,
		Key:    aws.String(key),
	})
	return err
}

func ObjectsSize(listObjectsInput *s3.ListObjectsV2Input) (int64, error) {
	res, err := client.ListObjectsV2(listObjectsInput)
	size := int64(0)
	for _, object := range res.Contents {
		size += *object.Size
	}
	return size, err
}

func BucketSize() (int64, error) {
	return ObjectsSize(&s3.ListObjectsV2Input{Bucket: bucket})
}

func BucketPrefixSize(prefix string) (int64, error) {
	return ObjectsSize(&s3.ListObjectsV2Input{
		Bucket: bucket,
		Prefix: aws.String(prefix),
	})
}
