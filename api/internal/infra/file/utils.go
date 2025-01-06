package file

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var qualities = map[string]int{"low": 10, "high": 90}

type result struct {
	Quality string
	Data    []byte
	Err     error
}

func compressImagesChan(ctx context.Context, img image.Image) <-chan result {
	var wg sync.WaitGroup
	results := make(chan result, len(qualities))

Loop:
	for lvl, quality := range qualities {
		select {
		case <-ctx.Done():
			break Loop
		default:
			wg.Add(1)
			go func(lvl string, quality int) {
				defer wg.Done()
				compressed, err := compressJPEG(img, quality)
				if err != nil {
					results <- result{Quality: lvl, Err: err}
					return
				}
				results <- result{Quality: lvl, Data: compressed}
			}(lvl, quality)
		}
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
