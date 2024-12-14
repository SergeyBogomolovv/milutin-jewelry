package usecase_test

import (
	"context"
	"mime/multipart"
	"testing"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/dto"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/entities"
	errs "github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/errors"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase"
	testutils "github.com/SergeyBogomolovv/milutin-jewelry/pkg/test-utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCollectionItemsUsecase_Create(t *testing.T) {
	mockRepo := new(collectionItemsRepoMock)
	mockFS := new(filesServiceMock)
	ctx := context.Background()
	usecase := usecase.NewCollectionItemsUsecase(testutils.NewTestLogger(), mockFS, mockRepo)
	t.Run("success", func(t *testing.T) {
		mockFS.On("UploadImage", ctx, mock.Anything, mock.Anything).Return("image_url", nil).Once()
		mockRepo.On("Create", ctx, mock.Anything).Return(1, nil).Once()

		id, err := usecase.Create(ctx, &dto.CreateCollectionItemRequest{Image: testutils.NewTestFile("image")})
		assert.NoError(t, err)
		assert.Equal(t, 1, id)
		mockFS.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})
	t.Run("failed to upload image", func(t *testing.T) {
		mockFS.On("UploadImage", ctx, mock.Anything, mock.Anything).Return("", assert.AnError).Once()
		_, err := usecase.Create(ctx, &dto.CreateCollectionItemRequest{Image: testutils.NewTestFile("image")})
		assert.Error(t, err)
		mockFS.AssertExpectations(t)
	})
	t.Run("collection not found", func(t *testing.T) {
		mockRepo.On("GetByID", ctx, mock.Anything).Return((*entities.Collection)(nil), errs.ErrCollectionNotFound).Once()
		_, err := usecase.Create(ctx, &dto.CreateCollectionItemRequest{CollectionID: 1})
		assert.ErrorIs(t, err, errs.ErrCollectionNotFound)
		mockRepo.AssertExpectations(t)
	})
	t.Run("failed to create collection-item", func(t *testing.T) {
		mockFS.On("UploadImage", ctx, mock.Anything, mock.Anything).Return("image_url", nil).Once()
		mockRepo.On("Create", ctx, mock.Anything).Return(0, assert.AnError).Once()
		_, err := usecase.Create(ctx, &dto.CreateCollectionItemRequest{Image: testutils.NewTestFile("image")})
		assert.Error(t, err)
	})
	mockFS.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestCollectionItemsUsecase_Update(t *testing.T) {
	mockRepo := new(collectionItemsRepoMock)
	mockFS := new(filesServiceMock)
	ctx := context.Background()
	usecase := usecase.NewCollectionItemsUsecase(testutils.NewTestLogger(), mockFS, mockRepo)
	t.Run("success", func(t *testing.T) {
		mockFS.On("UploadImage", ctx, mock.Anything, mock.Anything).Return("image_url", nil).Once()
		mockRepo.On("GetByID", ctx, mock.Anything).Return(&entities.CollectionItem{ImageID: "old_image_id"}, nil).Once()
		mockFS.On("DeleteImage", ctx, "old_image_id").Return(nil).Once()
		mockRepo.On("Update", ctx, mock.Anything).Return(nil).Once()

		err := usecase.Update(ctx, &dto.UpdateCollectionItemRequest{Image: testutils.NewTestFile("image")})

		assert.NoError(t, err)
		mockFS.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})
	t.Run("failed to upload image", func(t *testing.T) {
		mockFS.On("UploadImage", ctx, mock.Anything, mock.Anything).Return("", assert.AnError).Once()
		_, err := usecase.Create(ctx, &dto.CreateCollectionItemRequest{Image: testutils.NewTestFile("image")})
		assert.Error(t, err)
		mockFS.AssertExpectations(t)
	})
	t.Run("collection-item not found", func(t *testing.T) {
		mockRepo.On("GetByID", ctx, mock.Anything).Return((*entities.CollectionItem)(nil), errs.ErrCollectionItemNotFound).Once()
		err := usecase.Update(ctx, &dto.UpdateCollectionItemRequest{ID: 1})
		assert.ErrorIs(t, err, errs.ErrCollectionItemNotFound)
		mockRepo.AssertExpectations(t)
	})
	t.Run("failed to update collection-item", func(t *testing.T) {
		mockFS.On("UploadImage", ctx, mock.Anything, mock.Anything).Return("image_url", nil).Once()
		mockRepo.On("GetByID", ctx, mock.Anything).Return(&entities.CollectionItem{ImageID: "old_image_id"}, nil).Once()
		mockFS.On("DeleteImage", ctx, "old_image_id").Return(nil).Once()
		mockRepo.On("Update", ctx, mock.Anything).Return(assert.AnError).Once()
		err := usecase.Update(ctx, &dto.UpdateCollectionItemRequest{Image: testutils.NewTestFile("image")})
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
		mockFS.AssertExpectations(t)
	})
}

func TestCollectionItemsUsecase_Delete(t *testing.T) {
	mockRepo := new(collectionItemsRepoMock)
	mockFS := new(filesServiceMock)
	ctx := context.Background()
	usecase := usecase.NewCollectionItemsUsecase(testutils.NewTestLogger(), mockFS, mockRepo)
	t.Run("success", func(t *testing.T) {
		mockFS.On("DeleteImage", ctx, mock.Anything).Return(nil).Once()
		mockRepo.On("GetByID", ctx, 1).Return(&entities.CollectionItem{ImageID: "image_id", ID: 1}, nil).Once()
		mockRepo.On("Delete", ctx, 1).Return(nil).Once()

		err := usecase.Delete(ctx, 1)

		assert.NoError(t, err)
		mockFS.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})
	t.Run("collection-item not found", func(t *testing.T) {
		mockRepo.On("GetByID", ctx, 1).Return((*entities.CollectionItem)(nil), errs.ErrCollectionItemNotFound).Once()

		err := usecase.Delete(ctx, 1)

		assert.ErrorIs(t, err, errs.ErrCollectionItemNotFound)
		mockRepo.AssertExpectations(t)
	})
	t.Run("failed to delete collection-item", func(t *testing.T) {
		mockFS.On("DeleteImage", ctx, mock.Anything).Return(nil).Once()
		mockRepo.On("GetByID", ctx, 1).Return(&entities.CollectionItem{ImageID: "image_id", ID: 1}, nil).Once()
		mockRepo.On("Delete", ctx, 1).Return(assert.AnError).Once()
		err := usecase.Delete(ctx, 1)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
		mockFS.AssertExpectations(t)
	})
	t.Run("failed to delete image", func(t *testing.T) {
		mockFS.On("DeleteImage", ctx, mock.Anything).Return(assert.AnError).Once()
		mockRepo.On("GetByID", ctx, 1).Return(&entities.CollectionItem{ImageID: "image_id", ID: 1}, nil).Once()
		err := usecase.Delete(ctx, 1)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
		mockFS.AssertExpectations(t)
	})
}

func TestCollectionItemsUsecase_GetByCollection(t *testing.T) {
	mockRepo := new(collectionItemsRepoMock)
	ctx := context.Background()
	usecase := usecase.NewCollectionItemsUsecase(testutils.NewTestLogger(), nil, mockRepo)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetByCollection", ctx, 1).Return([]*entities.CollectionItem{{ID: 1}}, nil).Once()
		res, err := usecase.GetByCollection(ctx, 1)
		assert.NoError(t, err)
		assert.Len(t, res, 1)
		mockRepo.AssertExpectations(t)
	})
	t.Run("collection not found", func(t *testing.T) {
		mockRepo.On("GetByCollection", ctx, 1).Return(nil, errs.ErrCollectionNotFound).Once()
		_, err := usecase.GetByCollection(ctx, 1)
		assert.ErrorIs(t, err, errs.ErrCollectionNotFound)
		mockRepo.AssertExpectations(t)
	})
	t.Run("failed to get collection-items", func(t *testing.T) {
		mockRepo.On("GetByCollection", ctx, 1).Return(nil, assert.AnError).Once()
		_, err := usecase.GetByCollection(ctx, 1)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCollectionItemsUsecase_GetOne(t *testing.T) {
	mockRepo := new(collectionItemsRepoMock)
	ctx := context.Background()
	usecase := usecase.NewCollectionItemsUsecase(testutils.NewTestLogger(), nil, mockRepo)
	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetOne", ctx, 1).Return(&entities.CollectionItem{ID: 1}, nil).Once()
		res, err := usecase.GetOne(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, res.ID)
		mockRepo.AssertExpectations(t)
	})
	t.Run("collection-item not found", func(t *testing.T) {
		mockRepo.On("GetOne", ctx, 1).Return((*entities.CollectionItem)(nil), errs.ErrCollectionItemNotFound).Once()
		_, err := usecase.GetOne(ctx, 1)
		assert.ErrorIs(t, err, errs.ErrCollectionItemNotFound)
		mockRepo.AssertExpectations(t)
	})
	t.Run("failed to get collection-item", func(t *testing.T) {
		mockRepo.On("GetOne", ctx, 1).Return((*entities.CollectionItem)(nil), assert.AnError).Once()
		_, err := usecase.GetOne(ctx, 1)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

type collectionItemsRepoMock struct {
	mock.Mock
}

func (m *collectionItemsRepoMock) Create(ctx context.Context, dto *dto.CreateCollectionItemInput) (int, error) {
	args := m.Called(ctx, dto)
	return args.Int(0), args.Error(1)
}

func (m *collectionItemsRepoMock) Update(ctx context.Context, dto *dto.UpdateCollectionItemInput) error {
	args := m.Called(ctx, dto)
	return args.Error(0)
}

func (m *collectionItemsRepoMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *collectionItemsRepoMock) GetByCollection(ctx context.Context, id int) ([]*entities.CollectionItem, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]*entities.CollectionItem), args.Error(1)
}

func (m *collectionItemsRepoMock) GetOne(ctx context.Context, id int) (*entities.CollectionItem, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.CollectionItem), args.Error(1)
}

type filesServiceMock struct {
	mock.Mock
}

func (m *filesServiceMock) DeleteImage(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

func (m *filesServiceMock) UploadImage(ctx context.Context, file multipart.File, key string) (string, error) {
	args := m.Called(ctx, file, key)
	return args.String(0), args.Error(1)
}
