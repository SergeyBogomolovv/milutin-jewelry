package auth_test

import (
	"context"
	"testing"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/code"
	uc "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/auth"
	tu "github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/test"
	"github.com/stretchr/testify/assert"
)

func TestAuthUsecase_Login(t *testing.T) {
	ctx := context.Background()
	mockStorage := new(mockStorage)
	usecase := uc.New(tu.NewTestLogger(), mockStorage, nil, newTestConfig())

	t.Run("success", func(t *testing.T) {
		mockStorage.On("Check", ctx, "code").Return(nil).Once()
		mockStorage.On("Delete", ctx, "code").Return(nil).Once()
		token, err := usecase.Login(ctx, "code")
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		mockStorage.AssertExpectations(t)
	})

	t.Run("invalid code", func(t *testing.T) {
		mockStorage.On("Check", ctx, "code").Return(storage.ErrInvalidCode).Once()
		token, err := usecase.Login(ctx, "code")
		assert.ErrorIs(t, err, uc.ErrInvalidCode)
		assert.Empty(t, token)

		mockStorage.AssertExpectations(t)
	})

	t.Run("failed to check code", func(t *testing.T) {
		mockStorage.On("Check", ctx, "code").Return(assert.AnError).Once()
		token, err := usecase.Login(ctx, "code")
		assert.Error(t, err)
		assert.Empty(t, token)

		mockStorage.AssertExpectations(t)
	})

	t.Run("failed to delete code", func(t *testing.T) {
		mockStorage.On("Check", ctx, "code").Return(nil).Once()
		mockStorage.On("Delete", ctx, "code").Return(assert.AnError).Once()
		token, err := usecase.Login(ctx, "code")
		assert.Error(t, err)
		assert.Empty(t, token)

		mockStorage.AssertExpectations(t)
	})
}

func TestAuthUsecase_SendCode(t *testing.T) {
	ctx := context.Background()
	mockStorage := new(mockStorage)
	mockMailService := new(mockMailService)

	usecase := uc.New(tu.NewTestLogger(), mockStorage, mockMailService, newTestConfig())

	t.Run("success", func(t *testing.T) {
		mockStorage.On("Create", ctx).Return("code", nil).Once()
		mockMailService.On("SendCodeToAdmin", "code").Return(nil).Once()
		err := usecase.SendCode(ctx)
		assert.NoError(t, err)

		mockStorage.AssertExpectations(t)
	})

	t.Run("failed to create code", func(t *testing.T) {
		mockStorage.On("Create", ctx).Return("", assert.AnError).Once()
		err := usecase.SendCode(ctx)
		assert.Error(t, err)

		mockStorage.AssertExpectations(t)
	})

	t.Run("failed to send code", func(t *testing.T) {
		mockStorage.On("Create", ctx).Return("code", nil).Once()
		mockMailService.On("SendCodeToAdmin", "code").Return(assert.AnError).Once()
		err := usecase.SendCode(ctx)
		assert.Error(t, err)
		mockStorage.AssertExpectations(t)
		mockMailService.AssertExpectations(t)
	})
}
