package collection_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/collection"
	uc "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/collection"
	tu "github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/test"
)

func TestCollectionUsecase_Create(t *testing.T) {
	mockStorage := new(storageMock)
	mockFilesService := new(fileServiceMock)
	ctx := context.Background()
	usecase := uc.New(tu.NewTestLogger(), mockFilesService, mockStorage)

	t.Run("success", func(t *testing.T) {
		image := tu.NewTestFile("image")
		payload := uc.CreateCollectionPayload{Title: "test", Description: "description"}
		mockFilesService.On("UploadImage", ctx, image, mock.Anything).Return("image_id", nil).Once()
		mockStorage.On("Save", ctx, mock.Anything).Return(nil).Once()

		collection, err := usecase.Create(ctx, payload, image)
		assert.NoError(t, err)
		assert.Equal(t, payload.Title, collection.Title)
		mockFilesService.AssertExpectations(t)
		mockStorage.AssertExpectations(t)
	})

	t.Run("failed to upload image", func(t *testing.T) {
		image := tu.NewTestFile("image")
		payload := uc.CreateCollectionPayload{Title: "test"}
		mockFilesService.On("UploadImage", ctx, image, mock.Anything).Return("", assert.AnError).Once()

		collection, err := usecase.Create(ctx, payload, image)
		assert.Nil(t, collection)
		assert.Error(t, err)
		mockFilesService.AssertExpectations(t)
	})
}

func TestCollectionUsecase_Update(t *testing.T) {
	mockStorage := new(storageMock)
	mockFilesService := new(fileServiceMock)
	ctx := context.Background()
	usecase := uc.New(tu.NewTestLogger(), mockFilesService, mockStorage)

	t.Run("success", func(t *testing.T) {
		image := tu.NewTestFile("image")
		payload := uc.UpdateCollectionPayload{ID: 1, Title: ptr("new title")}
		stored := &storage.Collection{ID: 1, ImageID: "old_id", Title: "old title"}
		mockStorage.On("GetByID", ctx, payload.ID).Return(stored, nil).Once()
		mockFilesService.On("UploadImage", ctx, image, mock.Anything).Return("new_id", nil).Once()
		mockStorage.On("Update", ctx, mock.Anything).Return(nil).Once()
		mockFilesService.On("DeleteImage", ctx, stored.ImageID).Return(nil).Once()

		collection, err := usecase.Update(ctx, payload, image)
		assert.NoError(t, err)
		assert.Equal(t, "new_id", collection.ImageID)
		assert.Equal(t, *payload.Title, collection.Title)
		mockFilesService.AssertExpectations(t)
		mockStorage.AssertExpectations(t)
	})

	t.Run("failed to upload image", func(t *testing.T) {
		image := tu.NewTestFile("image")
		payload := uc.UpdateCollectionPayload{ID: 1}
		stored := &storage.Collection{ID: 1, Title: "old title"}
		mockStorage.On("GetByID", ctx, payload.ID).Return(stored, nil).Once()
		mockFilesService.On("UploadImage", ctx, image, mock.Anything).Return("", assert.AnError).Once()

		collection, err := usecase.Update(ctx, payload, image)
		assert.Nil(t, collection)
		assert.Error(t, err)
		mockFilesService.AssertExpectations(t)
	})
}

func TestCollectionUsecase_Delete(t *testing.T) {
	mockStorage := new(storageMock)
	mockFilesService := new(fileServiceMock)
	ctx := context.Background()
	usecase := uc.New(tu.NewTestLogger(), mockFilesService, mockStorage)

	t.Run("success", func(t *testing.T) {
		const id = 1
		stored := &storage.Collection{ID: id, ImageID: "image_id"}
		mockStorage.On("GetByID", ctx, id).Return(stored, nil).Once()
		mockFilesService.On("DeleteImage", ctx, stored.ImageID).Return(nil).Once()
		mockStorage.On("Delete", ctx, id).Return(nil).Once()

		collection, err := usecase.Delete(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, stored, collection)
		mockFilesService.AssertExpectations(t)
		mockStorage.AssertExpectations(t)
	})

	t.Run("failed to delete image", func(t *testing.T) {
		const id = 1
		stored := &storage.Collection{ID: id, ImageID: "image_id"}
		mockStorage.On("GetByID", ctx, id).Return(stored, nil).Once()
		mockStorage.On("Delete", ctx, id).Return(nil).Once()
		mockFilesService.On("DeleteImage", ctx, stored.ImageID).Return(assert.AnError).Once()

		collection, err := usecase.Delete(ctx, id)
		assert.Nil(t, collection)
		assert.Error(t, err)
		mockFilesService.AssertExpectations(t)
	})
}

func ptr(s string) *string {
	return &s
}
