package itemsusecase_test

import (
	"context"
	"testing"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/items"
	uc "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/items"
	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestItemsUsecase_Create(t *testing.T) {
	mockStorage := new(storageMock)
	mockFilesService := new(filesServiceMock)
	ctx := context.Background()
	usecase := uc.New(utils.NewTestLogger(), mockFilesService, mockStorage)

	t.Run("success", func(t *testing.T) {
		image := utils.NewTestFile("image")
		payload := uc.CreateItemPayload{CollectionID: 1, Title: "title"}
		mockStorage.On("CollectionExists", ctx, payload.CollectionID).Return(true, nil).Once()
		mockFilesService.On("UploadImage", ctx, image, mock.Anything).Return("image_id", nil).Once()
		mockStorage.On("Save", ctx, mock.Anything).Return(nil).Once()

		item, err := usecase.Create(ctx, payload, image)
		assert.NoError(t, err)
		assert.Equal(t, payload.Title, item.Title)
		mockFilesService.AssertExpectations(t)
		mockStorage.AssertExpectations(t)
	})

	t.Run("collection not found", func(t *testing.T) {
		payload := uc.CreateItemPayload{CollectionID: 1}
		mockStorage.On("CollectionExists", ctx, payload.CollectionID).Return(false, nil).Once()
		item, err := usecase.Create(ctx, payload, nil)
		assert.Nil(t, item)
		assert.ErrorIs(t, err, uc.ErrCollectionNotFound)
		mockStorage.AssertExpectations(t)
	})

	t.Run("failed to upload image", func(t *testing.T) {
		payload := uc.CreateItemPayload{CollectionID: 1}
		image := utils.NewTestFile("image")
		mockStorage.On("CollectionExists", ctx, payload.CollectionID).Return(true, nil).Once()
		mockFilesService.On("UploadImage", ctx, image, mock.Anything).Return("", assert.AnError).Once()
		item, err := usecase.Create(ctx, payload, image)
		assert.Nil(t, item)
		assert.Error(t, err)
		mockFilesService.AssertExpectations(t)
	})

	t.Run("storage error", func(t *testing.T) {
		image := utils.NewTestFile("image")
		payload := uc.CreateItemPayload{CollectionID: 1}
		mockStorage.On("CollectionExists", ctx, payload.CollectionID).Return(true, nil).Once()
		mockFilesService.On("UploadImage", ctx, image, mock.Anything).Return("image_id", nil).Once()
		mockStorage.On("Save", ctx, mock.Anything).Return(assert.AnError).Once()
		item, err := usecase.Create(ctx, payload, image)
		assert.Nil(t, item)
		assert.Error(t, err)
		mockFilesService.AssertExpectations(t)
		mockStorage.AssertExpectations(t)
	})
}

func TestItemsUsecase_Update(t *testing.T) {
	mockStorage := new(storageMock)
	mockFilesService := new(filesServiceMock)
	ctx := context.Background()
	usecase := uc.New(utils.NewTestLogger(), mockFilesService, mockStorage)

	t.Run("success", func(t *testing.T) {
		image := utils.NewTestFile("image")
		payload := uc.UpdateItemPayload{ID: 1}
		saved := &storage.Item{ImageID: "old_id"}
		mockStorage.On("GetById", ctx, payload.ID).Return(saved, nil).Once()
		mockFilesService.On("UploadImage", ctx, image, mock.Anything).Return("new_id", nil).Once()
		mockStorage.On("Update", ctx, mock.Anything).Return(nil).Once()
		mockFilesService.On("DeleteImage", ctx, saved.ImageID).Return(nil).Once()

		err := usecase.Update(ctx, payload, image)
		assert.NoError(t, err)
		mockFilesService.AssertExpectations(t)
		mockStorage.AssertExpectations(t)
	})

	t.Run("item not found", func(t *testing.T) {
		payload := uc.UpdateItemPayload{ID: 1}

		mockStorage.On("GetById", ctx, payload.ID).Return((*storage.Item)(nil), storage.ErrItemNotFound).Once()

		err := usecase.Update(ctx, payload, nil)
		assert.ErrorIs(t, err, uc.ErrItemNotFound)
		mockStorage.AssertExpectations(t)
	})

	t.Run("failed to upload image", func(t *testing.T) {
		image := utils.NewTestFile("image")
		payload := uc.UpdateItemPayload{ID: 1}

		mockStorage.On("GetById", ctx, payload.ID).Return(&storage.Item{}, nil).Once()
		mockFilesService.On("UploadImage", ctx, image, mock.Anything).Return("", assert.AnError).Once()

		err := usecase.Update(ctx, payload, image)
		assert.Error(t, err)
		mockFilesService.AssertExpectations(t)
	})

	t.Run("failed to delete image", func(t *testing.T) {
		image := utils.NewTestFile("image")
		payload := uc.UpdateItemPayload{ID: 1}
		saved := &storage.Item{ImageID: "image_id"}

		mockStorage.On("GetById", ctx, payload.ID).Return(saved, nil).Once()
		mockFilesService.On("UploadImage", ctx, image, mock.Anything).Return("new_id", nil).Once()
		mockStorage.On("Update", ctx, mock.Anything).Return(nil).Once()
		mockFilesService.On("DeleteImage", ctx, saved.ImageID).Return(assert.AnError).Once()

		err := usecase.Update(ctx, payload, image)
		assert.Error(t, err)
		mockFilesService.AssertExpectations(t)
		mockStorage.AssertExpectations(t)
	})

	t.Run("storage error", func(t *testing.T) {
		image := utils.NewTestFile("image")
		payload := uc.UpdateItemPayload{ID: 1}

		mockStorage.On("GetById", ctx, payload.ID).Return(&storage.Item{}, nil).Once()
		mockFilesService.On("UploadImage", ctx, image, mock.Anything).Return("image_id", nil).Once()
		mockStorage.On("Update", ctx, mock.Anything).Return(assert.AnError).Once()

		err := usecase.Update(ctx, payload, image)
		assert.Error(t, err)
		mockFilesService.AssertExpectations(t)
		mockStorage.AssertExpectations(t)
	})
}

func TestItemsUsecase_Delete(t *testing.T) {
	mockStorage := new(storageMock)
	mockFilesService := new(filesServiceMock)
	ctx := context.Background()
	usecase := uc.New(utils.NewTestLogger(), mockFilesService, mockStorage)

	t.Run("success", func(t *testing.T) {
		const id = 1
		saved := &storage.Item{ImageID: "image_id"}
		mockStorage.On("GetById", ctx, id).Return(saved, nil).Once()
		mockFilesService.On("DeleteImage", ctx, saved.ImageID).Return(nil).Once()
		mockStorage.On("Delete", ctx, id).Return(nil).Once()

		err := usecase.Delete(ctx, id)
		assert.NoError(t, err)
		mockFilesService.AssertExpectations(t)
		mockStorage.AssertExpectations(t)
	})

	t.Run("item not found", func(t *testing.T) {
		const id = 1
		mockStorage.On("GetById", ctx, id).Return((*storage.Item)(nil), storage.ErrItemNotFound).Once()
		err := usecase.Delete(ctx, id)
		assert.ErrorIs(t, err, uc.ErrItemNotFound)
		mockStorage.AssertExpectations(t)
	})

	t.Run("storage error", func(t *testing.T) {
		const id = 1
		mockStorage.On("GetById", ctx, id).Return(&storage.Item{}, nil).Once()
		mockStorage.On("Delete", ctx, id).Return(assert.AnError).Once()
		err := usecase.Delete(ctx, id)
		assert.Error(t, err)
		mockFilesService.AssertExpectations(t)
		mockStorage.AssertExpectations(t)
	})

	t.Run("failed to delete image", func(t *testing.T) {
		const id = 1
		saved := &storage.Item{ImageID: "image_id"}
		mockStorage.On("GetById", ctx, id).Return(saved, nil).Once()
		mockFilesService.On("DeleteImage", ctx, saved.ImageID).Return(assert.AnError).Once()
		mockStorage.On("Delete", ctx, id).Return(nil).Once()

		err := usecase.Delete(ctx, id)
		assert.Error(t, err)
		mockFilesService.AssertExpectations(t)
		mockStorage.AssertExpectations(t)
	})
}

func TestItemsUsecase_GetByCollection(t *testing.T) {
	mockStorage := new(storageMock)
	ctx := context.Background()
	usecase := uc.New(utils.NewTestLogger(), nil, mockStorage)

	t.Run("success", func(t *testing.T) {
		const collectionID = 1
		mockStorage.On("CollectionExists", ctx, collectionID).Return(true, nil).Once()
		mockStorage.On("GetByCollectionId", ctx, collectionID).Return([]*storage.Item{{ID: 1}, {ID: 2}}, nil).Once()
		items, err := usecase.GetByCollectionId(ctx, collectionID)
		assert.NoError(t, err)
		assert.Len(t, items, 2)
		mockStorage.AssertExpectations(t)
	})

	t.Run("collection not found", func(t *testing.T) {
		const collectionID = 1
		mockStorage.On("CollectionExists", ctx, collectionID).Return(false, nil).Once()
		_, err := usecase.GetByCollectionId(ctx, collectionID)
		assert.ErrorIs(t, err, uc.ErrCollectionNotFound)
		mockStorage.AssertExpectations(t)
	})

	t.Run("storage error", func(t *testing.T) {
		const collectionID = 1
		mockStorage.On("CollectionExists", ctx, collectionID).Return(true, nil).Once()
		mockStorage.On("GetByCollectionId", ctx, collectionID).Return(([]*storage.Item)(nil), assert.AnError).Once()
		res, err := usecase.GetByCollectionId(ctx, collectionID)
		assert.Error(t, err)
		assert.Nil(t, res)
		mockStorage.AssertExpectations(t)
	})
}
