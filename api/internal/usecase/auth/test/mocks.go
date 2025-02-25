package auth_test

import (
	"context"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/config"
	"github.com/stretchr/testify/mock"
)

type mockStorage struct {
	mock.Mock
}

func (m *mockStorage) Create(ctx context.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *mockStorage) Check(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}

func (m *mockStorage) Delete(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}

type mockMailService struct {
	mock.Mock
}

func (m *mockMailService) SendCodeToAdmin(code string) error {
	args := m.Called(code)
	return args.Error(0)
}

func newTestConfig() config.JwtConfig {
	return config.JwtConfig{
		Secret: []byte("secret"),
		TTL:    1,
	}
}
