package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	Endpoint   string `env:"S3_ENDPOINT_URL"`
	AccessKey  string `env:"S3_ACCESS_KEY"`
	SecretKey  string `env:"S3_SECRET_KEY"`
	Region     string `env:"S3_REGION"`
	BucketName string `env:"S3_BUCKET_NAME"`
}

var Cfg ServerConfig

func LoadConfig() {
	err := cleanenv.ReadConfig(".env", &Cfg)
	if err != nil {
		log.Fatal(err)
	}
}
