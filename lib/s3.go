package lib

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3config "github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func getS3Client() *s3.Client {
	cfg, err := s3config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err.Error())
	}
	return s3.NewFromConfig(cfg)
}

func UploadS3(file multipart.File, filename string) (string, error) {
	client := getS3Client()

	uploader := manager.NewUploader(client)

	//* Old way
	// result, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
	// 	Bucket: aws.String("go-fiber-upload-test"),
	// 	Key:    aws.String(filename),
	// 	Body:   file,
	// })

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("go-fiber-upload-test-v2"),
		Key:    aws.String(filename),
		//ACL:    "public-read", <- Ignore this
		Body: file,
	})

	if err != nil {
		return "", err
	}

	return result.Location, nil
}

func GetFilesS3() ([]types.Object, error) {
	client := getS3Client()

	result, err := client.ListObjects(context.TODO(), &s3.ListObjectsInput{
		Bucket: aws.String("go-fiber-upload-test"),
	})

	if err != nil {
		return nil, err
	}

	return result.Contents, nil
}

func GetFileS3(key string) (string, error) {
	client := s3.NewPresignClient(getS3Client())

	result, err := client.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String("go-fiber-upload-test"),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(15*time.Minute))

	if err != nil {
		return "", err
	}

	return result.URL, nil
}
