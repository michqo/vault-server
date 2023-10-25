package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	Endpoint           string `env:"S3_ENDPOINT_URL" env-required:"true"`
	AccessKey          string `env:"S3_ACCESS_KEY" env-required:"true"`
	SecretKey          string `env:"S3_SECRET_KEY" env-required:"true"`
	Region             string `env:"S3_REGION" env-required:"true"`
	BucketName         string `env:"S3_BUCKET_NAME" env-required:"true"`
	MaxFolderSize      int64  `env:"MAX_FOLDER_SIZE" env-default:"30"`
	MaxBucketSize      int64  `env:"MAX_BUCKET_SIZE" env-default:"10"`
	MaxObjectFetchSize int64  `env:"MAX_OBJECT_FETCH_SIZE" env-default:"4"`
}

const MB int64 = 1048576

var Cfg ServerConfig

func LoadConfig() {
	err := cleanenv.ReadConfig(".env", &Cfg)
	if err != nil {
		log.Fatal(err)
	}
	Cfg.MaxFolderSize *= MB
	Cfg.MaxBucketSize *= MB * 1000
	Cfg.MaxObjectFetchSize *= MB
}
