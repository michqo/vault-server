package database

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var client *s3.S3
var bucket *string

func CreateDatabase() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("S3_REGION")),
		Endpoint:    aws.String(os.Getenv("S3_ENDPOINT_URL")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_ACCESS_KEY"), os.Getenv("S3_SECRET_KEY"), ""),
	})
	if err != nil {
		log.Fatal(err)
	}
	bucket = aws.String(os.Getenv("S3_BUCKET_NAME"))
	client = s3.New(sess)
}

func ListObjects(prefix string) (*s3.ListObjectsV2Output, error) {
	res, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: bucket,
		Prefix: aws.String(prefix),
	})
	return res, err
}

func GetObjectUrl(key string) (string, error) {
	res, _ := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket:                     bucket,
		Key:                        aws.String(key),
		ResponseContentDisposition: aws.String("attachment"),
	})
	url, err := res.Presign(60 * time.Minute)
	return url, err
}

func PutObjectUrl(key string) (string, error) {
	res, _ := client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: bucket,
		Key:    aws.String(key),
	})
	url, err := res.Presign(5 * time.Minute)
	return url, err
}

func DeleteObject(key string) error {
	_, err := client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: bucket,
		Key:    aws.String(key),
	})
	return err
}
