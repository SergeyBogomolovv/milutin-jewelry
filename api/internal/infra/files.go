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

func (f *filesService) Upload(ctx context.Context, key string, data []byte) error {
	_, err := f.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(f.bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})
	return err
}

func (s *filesService) UploadImage(ctx context.Context, header *multipart.FileHeader, key string) (string, error) {
	s.log.Info("uploading image", "key", key)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	file, err := header.Open()
	if err != nil {
		s.log.Error("failed to open file", "err", err)
		return "", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		s.log.Error("failed to decode image", "err", err)
		return "", err
	}
	imageID := uuid.NewString()

	results := compressImage(ctx, img)

	for res := range results {
		if res.Err != nil {
			cancel()
			s.log.Error("failed to compress image", "err", res.Err, "quality", res.Quality)
			return "", res.Err
		}
		if err := s.Upload(ctx, fmt.Sprintf("%s/%s_%s.jpg", key, imageID, res.Quality), res.Data); err != nil {
			cancel()
			return "", err
		}
	}
	return imageID, nil
}

type result struct {
	Quality string
	Data    []byte
	Err     error
}

func compressImage(ctx context.Context, img image.Image) <-chan result {
	var wg sync.WaitGroup
	qualities := map[string]int{"low": 30, "med": 60, "high": 90}
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

func compressJPEG(img image.Image, quality int) ([]byte, error) {
	var buff bytes.Buffer
	if err := jpeg.Encode(&buff, img, &jpeg.Options{Quality: quality}); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
