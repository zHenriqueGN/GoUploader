package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

type envVars struct {
	NumFiles     int    `mapstructure:"NUM_FILES"`
	S3BucketName string `mapstructure:"S3_BUCKET_NAME"`
	AWSRegion    string `mapstructure:"AWS_REGION"`
	AWSID        string `mapstructure:"AWS_ID"`
	AWSSecret    string `mapstructure:"AWS_SECRET"`
	AWSToken     string `mapstructure:"AWS_TOKEN"`
	S3Client     *s3.S3
}

var EnvVars envVars

func LoadConfigs() *envVars {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&EnvVars)
	if err != nil {
		panic(err)
	}
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(EnvVars.AWSRegion),
			Credentials: credentials.NewStaticCredentials(
				EnvVars.AWSID,
				EnvVars.AWSSecret,
				EnvVars.AWSToken,
			),
		},
	)
	if err != nil {
		panic(err)
	}
	EnvVars.S3Client = s3.New(sess)
	return &EnvVars
}
