package uploader

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/zHenriqueGN/GoUploader/internal/config"
)

func UploadFile(fileName string) {
	filePath := fmt.Sprintf("./tmp/%s", fileName)
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()
	S3Client := config.EnvVars.S3Client
	_, err = S3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(config.EnvVars.S3BucketName),
		Key:    aws.String(config.EnvVars.AWSID),
		Body:   file,
	})
	if err != nil {
		log.Printf("error uploading file %s: %v\n", filePath, err)
		return
	}
	fmt.Printf("file %s uploaded successfully\n", filePath)
}
