package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

type EnvVars struct {
	NumFiles     int    `mapstructure:"NUM_FILES"`
	S3BucketName string `mapstructure:"S3_BUCKET_NAME"`
	AWSRegion    string `mapstructure:"AWS_REGION"`
	AWSID        string `mapstructure:"AWS_ID"`
	AWSSecret    string `mapstructure:"AWS_SECRET"`
	AWSToken     string `mapstructure:"AWS_TOKEN"`
	S3Client     *s3.S3
}

func LoadConfigs() *EnvVars {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var envVars EnvVars
	err = viper.Unmarshal(&envVars)
	if err != nil {
		panic(err)
	}
	sess, err := session.NewSession(
		&aws.Config{
			Region: &envVars.AWSRegion,
			Credentials: credentials.NewStaticCredentials(
				envVars.AWSID,
				envVars.AWSSecret,
				envVars.AWSToken,
			),
		},
	)
	if err != nil {
		panic(err)
	}
	envVars.S3Client = s3.New(sess)
	return &envVars
}
