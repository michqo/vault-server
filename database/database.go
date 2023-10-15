package database

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var client *s3.S3
var bucket string

func CreateDatabase() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("S3_REGION")),
		Endpoint:    aws.String(os.Getenv("S3_ENDPOINT_URL")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("S3_ACCESS_KEY"), os.Getenv("S3_SECRET_KEY"), ""),
	})
	if err != nil {
		log.Fatal(err)
	}
	bucket = os.Getenv("S3_BUCKET_NAME")
	client = s3.New(sess)
}

func ListItems(prefix string) (*s3.ListObjectsV2Output, error) {
	res, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
