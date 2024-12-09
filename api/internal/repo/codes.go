package repo

import (
	"context"
	"crypto/rand"
	"math/big"
	"time"

	"github.com/redis/go-redis/v9"
)

type codesRepo struct {
	redis *redis.Client
}

func NewCodesRepo(redis *redis.Client) *codesRepo {
	return &codesRepo{redis: redis}
}

func (r *codesRepo) CheckCode(ctx context.Context, code string) error {
	return r.redis.Get(ctx, code).Err()
}

func (r *codesRepo) CreateCode(ctx context.Context) (string, error) {
	code, err := generateCode()
	if err != nil {
		return "", err
	}
	return code, r.redis.Set(ctx, code, true, time.Minute*5).Err()
}

func (r *codesRepo) DeleteCode(ctx context.Context, code string) error {
	return r.redis.Del(ctx, code).Err()
}

func generateCode() (string, error) {
	const digits = "0123456789"

	otp := make([]byte, 6)
	for i := range otp {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[num.Int64()]
	}

	return string(otp), nil
}
