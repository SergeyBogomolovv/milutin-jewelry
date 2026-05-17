package code

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
	storedCode, err := r.redis.Get(ctx, codeKey).Result()
	if err != nil {
		return ErrInvalidCode
	}
	if storedCode != code {
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
	if err := r.redis.Set(ctx, codeKey, code, time.Minute*5).Err(); err != nil {
		return "", err
	}
	return code, nil
}

func (r *storage) Delete(ctx context.Context, code string) error {
	const deleteIfCodeMatches = `
if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
end
return 0`
	deleted, err := r.redis.Eval(ctx, deleteIfCodeMatches, []string{codeKey}, code).Int()
	if err != nil {
		return ErrCodeNotFound
	}
	if deleted == 0 {
		return ErrCodeNotFound
	}
	return nil
}
