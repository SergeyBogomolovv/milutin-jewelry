package files

import (
	"context"
	"fmt"
	"image"
	_ "image/png"
	"log/slog"
	"mime/multipart"
	"sync"

	cfg "github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
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

func (s *filesService) DeleteImage(ctx context.Context, key string) error {
	const op = "DeleteImage"
	log := s.log.With(slog.String("op", op), slog.String("key", key))

	log.Info("deleting image")
	var wg sync.WaitGroup

	for quality := range qualities {
		wg.Add(1)
		go func(quality string) {
			defer wg.Done()
			if err := s.delete(ctx, fmt.Sprintf("%s_%s.jpg", key, quality)); err != nil {
				log.Error("failed to delete image", "err", err, "quality", quality)
			}
		}(quality)
	}
	wg.Wait()
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

	results := compressImagesChan(ctx, img)

	for res := range results {
		if res.Err != nil {
			cancel()
			log.Error("failed to compress image", "err", res.Err, "quality", res.Quality)
			return "", res.Err
		}
		if err := s.upload(ctx, fmt.Sprintf("%s/%s_%s.jpg", path, imageID, res.Quality), res.Data); err != nil {
			cancel()
			log.Error("failed to upload image", "err", err, "quality", res.Quality)
			return "", err
		}
	}

	return fmt.Sprintf("%s/%s", path, imageID), nil
}
