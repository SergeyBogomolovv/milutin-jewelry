package codestorage

import (
	"context"
	"time"

	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/e"
	"github.com/redis/go-redis/v9"
)

type storage struct {
	redis *redis.Client
}

func New(redis *redis.Client) *storage {
	return &storage{redis: redis}
}

func (r *storage) Check(ctx context.Context, code string) error {
	if err := r.redis.Get(ctx, codeString(code)).Err(); err != nil {
		return ErrInvalidCode
	}
	return nil
}

func (r *storage) Create(ctx context.Context) (code string, err error) {
	defer func() { err = e.WrapIfErr("can't create code", err) }()
	code, err = generateCode()
	if err != nil {
		return "", err
	}
	if err := r.redis.Set(ctx, codeString(code), true, time.Minute*5).Err(); err != nil {
		return "", err
	}
	return code, nil
}

func (r *storage) Delete(ctx context.Context, code string) error {
	if err := r.redis.Del(ctx, codeString(code)).Err(); err != nil {
		return ErrCodeNotFound
	}
	return nil
}
