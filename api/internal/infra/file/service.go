package file

import (
	"context"
	"fmt"
	"image"
	_ "image/png"
	"log/slog"
	"mime/multipart"

	cfg "github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/e"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type filesService struct {
	client *s3.Client
	bucket string
	log    *slog.Logger
}

func New(log *slog.Logger, c cfg.ObjectStorageConfig) *filesService {
	const dest = "filesService"
	config, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(c.AccessKey, c.SecretKey, "")),
		config.WithRegion(c.Region),
		config.WithBaseEndpoint(c.Endpoint),
	)
	if err != nil {
		panic(err)
	}
	client := s3.NewFromConfig(config)

	return &filesService{client: client, bucket: c.Bucket, log: log.With(slog.String("dest", dest))}
}

func (s *filesService) DeleteImage(ctx context.Context, key string) (err error) {
	defer func() { err = e.WrapIfErr("can't delete image", err) }()
	const op = "DeleteImage"
	log := s.log.With(slog.String("op", op), slog.String("key", key))

	log.Info("deleting image")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := s.delete(ctx, fmt.Sprintf("%s.jpg", key)); err != nil {
		log.Error("failed to delete image", "err", err)
		return err
	}
	if err := s.delete(ctx, fmt.Sprintf("%s_low.jpg", key)); err != nil {
		log.Error("failed to delete image", "err", err)
		return err
	}
	return nil
}

func (s *filesService) UploadImage(ctx context.Context, file multipart.File, path string) (string, error) {
	const op = "UploadImage"
	log := s.log.With(slog.String("op", op), slog.String("path", path))

	log.Info("uploading image")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Error("failed to decode image", "err", err)
		return "", err
	}
	imageID := uuid.NewString()

	compressedHigh, err := compressHigh(img, 100)
	if err != nil {
		log.Error("failed to compress image to jpeg", "err", err)
		return "", err
	}

	compressedLow, err := compressLow(img, 10)
	if err != nil {
		log.Error("failed to compress image to base64", "err", err)
		return "", err
	}

	if err := s.upload(ctx, fmt.Sprintf("%s/%s.jpg", path, imageID), compressedHigh); err != nil {
		log.Error("failed to upload image", "err", err)
		return "", err
	}

	if err := s.upload(ctx, fmt.Sprintf("%s/%s_low.jpg", path, imageID), compressedLow); err != nil {
		log.Error("failed to upload image", "err", err)
		return "", err
	}

	return fmt.Sprintf("%s/%s", path, imageID), nil
}
