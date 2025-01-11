package file

import (
	"bytes"
	"context"
	"fmt"
	"image"

	"golang.org/x/image/draw"

	"image/jpeg"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func lowKey(key string) string {
	return fmt.Sprintf("%s_low.jpg", key)
}

func highKey(key string) string {
	return fmt.Sprintf("%s.jpg", key)
}

func compressHigh(img image.Image, quality int) ([]byte, error) {
	var buff bytes.Buffer
	if err := jpeg.Encode(&buff, img, &jpeg.Options{Quality: quality}); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func compressLow(img image.Image, bound int) ([]byte, error) {
	dst := image.NewRGBA(image.Rect(0, 0, bound, bound))
	draw.NearestNeighbor.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)

	var buff bytes.Buffer
	err := jpeg.Encode(&buff, dst, nil)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
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
