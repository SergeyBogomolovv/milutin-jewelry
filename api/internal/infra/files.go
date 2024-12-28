package infra

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"log/slog"
	"mime/multipart"
	"sync"

	cfg "github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
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

var qualities = map[string]int{"low": 1, "high": 90}

func NewFilesService(log *slog.Logger, c cfg.ObjectStorageConfig) *filesService {
	config, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(c.AccessKey, c.SecretKey, "")),
		config.WithRegion(c.Region),
		config.WithBaseEndpoint(c.Endpoint),
	)
	if err != nil {
		panic(err)
	}
	client := s3.NewFromConfig(config)

	return &filesService{client: client, bucket: c.Bucket, log: log.With(slog.String("op", "filesService"))}
}

func (s *filesService) DeleteImage(ctx context.Context, key string) error {
	s.log.Info("deleting image", "key", key)
	var wg sync.WaitGroup

	for quality := range qualities {
		wg.Add(1)
		go func(quality string) {
			defer wg.Done()
			if err := s.delete(ctx, fmt.Sprintf("%s_%s.jpg", key, quality)); err != nil {
				s.log.Error("failed to delete image", "err", err, "key", key, "quality", quality)
			}
		}(quality)
	}
	wg.Wait()
	return nil
}

func (s *filesService) UploadImage(ctx context.Context, file multipart.File, key string) (string, error) {
	s.log.Info("uploading image", "key", key)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	img, _, err := image.Decode(file)
	if err != nil {
		s.log.Error("failed to decode image", "err", err)
		return "", err
	}
	imageID := uuid.NewString()

	results := compressImagesChan(ctx, img)

	for res := range results {
		if res.Err != nil {
			cancel()
			s.log.Error("failed to compress image", "err", res.Err, "quality", res.Quality)
			return "", res.Err
		}
		if err := s.upload(ctx, fmt.Sprintf("%s/%s_%s.jpg", key, imageID, res.Quality), res.Data); err != nil {
			cancel()
			s.log.Error("failed to upload image", "err", err, "quality", res.Quality)
			return "", err
		}
	}

	return fmt.Sprintf("%s/%s", key, imageID), nil
}

type result struct {
	Quality string
	Data    []byte
	Err     error
}

func compressImagesChan(ctx context.Context, img image.Image) <-chan result {
	var wg sync.WaitGroup
	results := make(chan result, len(qualities))

	for lvl, quality := range qualities {
		if ctx.Err() != nil {
			break
		}
		wg.Add(1)
		go func(lvl string, quality int) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				compressed, err := compressJPEG(img, quality)
				if err != nil {
					results <- result{Quality: lvl, Err: err}
					return
				}
				results <- result{Quality: lvl, Data: compressed}
			}
		}(lvl, quality)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func (s *filesService) delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return err
}

func (f *filesService) upload(ctx context.Context, key string, data []byte) error {
	_, err := f.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(f.bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})
	return err
}

func compressJPEG(img image.Image, quality int) ([]byte, error) {
	var buff bytes.Buffer
	if err := jpeg.Encode(&buff, img, &jpeg.Options{Quality: quality}); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
