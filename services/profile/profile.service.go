package profile

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsS3 "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"goFiber/main/config"
	"goFiber/main/models"
	"log"
	"mime/multipart"
)

func CreateProfileService(newProfile models.Profile) (models.Profile, error) {
	config.Database.Db.Create(&newProfile)

	return newProfile, nil
}

func ImageFileService(file *multipart.FileHeader) (*manager.UploadOutput, error) {
	f, err := file.Open()

	if err != nil {
		log.Println("file upload error")
	}

	cfg, err := awsS3.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("sketchmind"),
		Key:    aws.String(file.Filename),
		Body:   f,
	})

	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return result, nil
}
