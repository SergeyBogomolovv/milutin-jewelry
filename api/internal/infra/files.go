package infra

import (
	"bytes"
	"context"

	cfg "github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type filesService struct {
	client *s3.Client
	bucket string
}

func NewFilesService(c cfg.ObjectStorageConfig) *filesService {
	config, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(c.AccessKey, c.SecretKey, "")),
		config.WithRegion(c.Region),
		config.WithBaseEndpoint(c.Endpoint),
	)
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(config)

	return &filesService{client: client, bucket: c.Bucket}
}

func (f *filesService) Upload(ctx context.Context, key string, data []byte) error {
	_, err := f.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(f.bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})
	return err
}
