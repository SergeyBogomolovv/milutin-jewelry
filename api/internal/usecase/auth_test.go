package usecase_test

import (
	"context"
	"testing"

	errs "github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/errors"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase"
	testutils "github.com/SergeyBogomolovv/milutin-jewelry/pkg/test-utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthUsecase_Login(t *testing.T) {
	ctx := context.Background()
	mockCodesRepo := new(mockCodesRepo)

	usecase := usecase.NewAuthUsecase(testutils.NewTestLogger(), mockCodesRepo, nil, "secret")

	t.Run("success", func(t *testing.T) {
		mockCodesRepo.On("Check", ctx, "code").Return(nil).Once()
		mockCodesRepo.On("Delete", ctx, "code").Return(nil).Once()
		token, err := usecase.Login(ctx, "code")
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		mockCodesRepo.AssertExpectations(t)
	})

	t.Run("invalid code", func(t *testing.T) {
		mockCodesRepo.On("Check", ctx, "code").Return(errs.ErrInvalidLoginCode).Once()
		token, err := usecase.Login(ctx, "code")
		assert.ErrorIs(t, err, errs.ErrInvalidLoginCode)
		assert.Empty(t, token)

		mockCodesRepo.AssertExpectations(t)
	})

	t.Run("failed to check code", func(t *testing.T) {
		mockCodesRepo.On("Check", ctx, "code").Return(assert.AnError).Once()
		token, err := usecase.Login(ctx, "code")
		assert.Error(t, err)
		assert.Empty(t, token)

		mockCodesRepo.AssertExpectations(t)
	})

	t.Run("failed to delete code", func(t *testing.T) {
		mockCodesRepo.On("Check", ctx, "code").Return(nil).Once()
		mockCodesRepo.On("Delete", ctx, "code").Return(assert.AnError).Once()
		token, err := usecase.Login(ctx, "code")
		assert.Error(t, err)
		assert.Empty(t, token)

		mockCodesRepo.AssertExpectations(t)
	})
}

func TestAuthUsecase_SendCode(t *testing.T) {
	ctx := context.Background()
	mockCodesRepo := new(mockCodesRepo)
	mockMailService := new(mockMailService)

	usecase := usecase.NewAuthUsecase(testutils.NewTestLogger(), mockCodesRepo, mockMailService, "secret")

	t.Run("success", func(t *testing.T) {
		mockCodesRepo.On("Create", ctx).Return("code", nil).Once()
		mockMailService.On("SendCodeToAdmin", ctx, "code").Return(nil).Once()
		err := usecase.SendCode(ctx)
		assert.NoError(t, err)

		mockCodesRepo.AssertExpectations(t)
	})

	t.Run("failed to create code", func(t *testing.T) {
		mockCodesRepo.On("Create", ctx).Return("", assert.AnError).Once()
		err := usecase.SendCode(ctx)
		assert.Error(t, err)

		mockCodesRepo.AssertExpectations(t)
	})

	t.Run("failed to send email", func(t *testing.T) {
		mockCodesRepo.On("Create", ctx).Return("code", nil).Once()
		mockMailService.On("SendCodeToAdmin", ctx, "code").Return(assert.AnError).Once()
		err := usecase.SendCode(ctx)
		assert.Error(t, err)

		mockCodesRepo.AssertExpectations(t)
	})
}

type mockCodesRepo struct {
	mock.Mock
}

func (m *mockCodesRepo) Create(ctx context.Context) (string, error) {
	args := m.Called(ctx)
	return args.String(0), args.Error(1)
}

func (m *mockCodesRepo) Check(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}

func (m *mockCodesRepo) Delete(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}

type mockMailService struct {
	mock.Mock
}

func (m *mockMailService) SendCodeToAdmin(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}
