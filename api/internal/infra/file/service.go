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
	"golang.org/x/sync/errgroup"
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

func (s *filesService) DeleteImage(ctx context.Context, key string) error {
	const op = "DeleteImage"
	log := s.log.With(slog.String("op", op), slog.String("key", key))

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return s.delete(ctx, lowKey(key))
	})
	eg.Go(func() error {
		return s.delete(ctx, highKey(key))
	})

	if err := eg.Wait(); err != nil {
		log.Error("failed to delete image", "err", err)
		return e.Wrap("failed to delete image", err)
	}

	log.Info("image deleted")

	return nil
}

func (s *filesService) UploadImage(ctx context.Context, file multipart.File, path string) (string, error) {
	const op = "UploadImage"
	log := s.log.With(slog.String("op", op), slog.String("path", path))

	img, _, err := image.Decode(file)
	if err != nil {
		log.Error("failed to decode image", "err", err)
		return "", err
	}
	imageID := uuid.NewString()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		data, err := compressHigh(img, 90)
		if err != nil {
			log.Error("failed to compress high image", "err", err)
			return err
		}
		return s.upload(ctx, highKey(fmt.Sprintf("%s/%s", path, imageID)), data)
	})

	eg.Go(func() error {
		data, err := compressLow(img, 50)
		if err != nil {
			log.Error("failed to compress low image", "err", err)
			return err
		}
		return s.upload(ctx, lowKey(fmt.Sprintf("%s/%s", path, imageID)), data)
	})

	if err := eg.Wait(); err != nil {
		log.Error("failed to upload image", "err", err)
		return "", err
	}

	key := fmt.Sprintf("%s/%s", path, imageID)
	log.Info("image uploaded", "key", key)
	return key, nil
}
