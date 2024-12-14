package usecase

import (
	"context"
	"mime/multipart"
	"testing"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/dto"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/entities"
	errs "github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/errors"
	testutils "github.com/SergeyBogomolovv/milutin-jewelry/pkg/test-utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCollectionsUsecase_CreateCollection(t *testing.T) {
	mockFS := new(filesServiceMock)
	mockRepo := new(collectionsRepoMock)
	ctx := context.Background()

	usecase := NewCollectionsUsecase(testutils.NewTestLogger(), mockFS, mockRepo)
	t.Run("success", func(t *testing.T) {
		mockFS.On("UploadImage", ctx, mock.Anything, mock.Anything).Return("image_url", nil).Once()
		mockRepo.On("CreateCollection", ctx, mock.Anything).Return(1, nil).Once()

		id, err := usecase.CreateCollection(ctx, &dto.CreateCollectionRequest{Image: testutils.NewTestFile("image")})
		assert.NoError(t, err)
		assert.Equal(t, 1, id)

		mockFS.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})
	t.Run("failed to upload image", func(t *testing.T) {
		mockFS.On("UploadImage", ctx, mock.Anything, mock.Anything).Return("", assert.AnError).Once()
		_, err := usecase.CreateCollection(ctx, &dto.CreateCollectionRequest{Image: testutils.NewTestFile("image")})
		assert.Error(t, err)
		mockFS.AssertExpectations(t)
	})
	t.Run("failed to create collection", func(t *testing.T) {
		mockFS.On("UploadImage", ctx, mock.Anything, mock.Anything).Return("image_url", nil).Once()
		mockRepo.On("CreateCollection", ctx, mock.Anything).Return(0, assert.AnError).Once()
		_, err := usecase.CreateCollection(ctx, &dto.CreateCollectionRequest{Image: testutils.NewTestFile("image")})
		assert.Error(t, err)
		mockFS.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})
}

func TestCollectionsUsecase_UpdateCollection(t *testing.T) {
	mockFS := new(filesServiceMock)
	mockRepo := new(collectionsRepoMock)
	ctx := context.Background()

	usecase := NewCollectionsUsecase(testutils.NewTestLogger(), mockFS, mockRepo)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetCollectionByID", ctx, mock.Anything).Return(&entities.Collection{ImageID: "old_image_id"}, nil).Once()
		mockFS.On("DeleteImage", ctx, "old_image_id").Return(nil).Once()
		mockFS.On("UploadImage", ctx, mock.Anything, "collections").Return("new_image_id", nil).Once()
		mockRepo.On("UpdateCollection", ctx, mock.Anything).Return(nil).Once()

		err := usecase.UpdateCollection(ctx, &dto.UpdateCollectionRequest{
			Image: testutils.NewTestFile("image"),
		})

		assert.NoError(t, err)
		mockFS.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})

	t.Run("collection not found", func(t *testing.T) {
		mockRepo.On("GetCollectionByID", ctx, mock.Anything).Return((*entities.Collection)(nil), errs.ErrCollectionNotFound).Once()
		err := usecase.UpdateCollection(ctx, &dto.UpdateCollectionRequest{ID: 1})

		assert.ErrorIs(t, err, errs.ErrCollectionNotFound)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed to delete image", func(t *testing.T) {
		mockRepo.On("GetCollectionByID", ctx, mock.Anything).Return(&entities.Collection{ImageID: "old_image_id"}, nil).Once()
		mockFS.On("DeleteImage", ctx, "old_image_id").Return(assert.AnError).Once()

		err := usecase.UpdateCollection(ctx, &dto.UpdateCollectionRequest{Image: testutils.NewTestFile("image")})

		assert.Error(t, err)
		mockFS.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed to upload image", func(t *testing.T) {
		mockRepo.On("GetCollectionByID", ctx, mock.Anything).Return(&entities.Collection{ImageID: "old_image_id"}, nil).Once()
		mockFS.On("DeleteImage", ctx, "old_image_id").Return(nil).Once()
		mockFS.On("UploadImage", ctx, mock.Anything, "collections").Return("", assert.AnError).Once()

		err := usecase.UpdateCollection(ctx, &dto.UpdateCollectionRequest{Image: testutils.NewTestFile("image")})

		assert.Error(t, err)
		mockFS.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed to update collection", func(t *testing.T) {
		mockRepo.On("GetCollectionByID", ctx, mock.Anything).Return(&entities.Collection{ImageID: "old_image_id"}, nil).Once()
		mockFS.On("DeleteImage", ctx, "old_image_id").Return(nil).Once()
		mockFS.On("UploadImage", ctx, mock.Anything, "collections").Return("new_image_id", nil).Once()
		mockRepo.On("UpdateCollection", ctx, mock.Anything).Return(assert.AnError).Once()

		err := usecase.UpdateCollection(ctx, &dto.UpdateCollectionRequest{Image: testutils.NewTestFile("image")})

		assert.Error(t, err)
		mockFS.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})
}

func TestCollectionsUsecase_GetCollectionByID(t *testing.T) {
	mockRepo := new(collectionsRepoMock)
	ctx := context.Background()

	usecase := NewCollectionsUsecase(testutils.NewTestLogger(), nil, mockRepo)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetCollectionByID", ctx, 1).Return(&entities.Collection{
			Title: "test",
		}, nil).Once()

		result, err := usecase.GetCollectionByID(ctx, 1)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "test", result.Title)
		mockRepo.AssertExpectations(t)
	})

	t.Run("collection not found", func(t *testing.T) {
		mockRepo.On("GetCollectionByID", ctx, 1).Return((*entities.Collection)(nil), errs.ErrCollectionNotFound).Once()

		result, err := usecase.GetCollectionByID(ctx, 1)

		assert.ErrorIs(t, err, errs.ErrCollectionNotFound)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed to get collection", func(t *testing.T) {
		mockRepo.On("GetCollectionByID", ctx, 1).Return((*entities.Collection)(nil), assert.AnError).Once()

		result, err := usecase.GetCollectionByID(ctx, 1)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestCollectionsUsecase_GetAllCollections(t *testing.T) {
	mockRepo := new(collectionsRepoMock)
	ctx := context.Background()

	usecase := NewCollectionsUsecase(testutils.NewTestLogger(), nil, mockRepo)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetAllCollections", ctx).Return([]*entities.Collection{{ID: 1}, {ID: 2}}, nil).Once()

		result, err := usecase.GetAllCollections(ctx)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed to get collections", func(t *testing.T) {
		mockRepo.On("GetAllCollections", ctx).Return(([]*entities.Collection)(nil), assert.AnError).Once()

		result, err := usecase.GetAllCollections(ctx)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestCollectionsUsecase_DeleteCollection(t *testing.T) {
	mockFS := new(filesServiceMock)
	mockRepo := new(collectionsRepoMock)
	ctx := context.Background()

	usecase := NewCollectionsUsecase(testutils.NewTestLogger(), mockFS, mockRepo)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetCollectionByID", ctx, 1).Return(&entities.Collection{
			ID:      1,
			ImageID: "image_id",
		}, nil).Once()
		mockFS.On("DeleteImage", ctx, "image_id").Return(nil).Once()
		mockRepo.On("DeleteCollection", ctx, 1).Return(nil).Once()

		err := usecase.DeleteCollection(ctx, 1)

		assert.NoError(t, err)
		mockFS.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})

	t.Run("collection not found", func(t *testing.T) {
		mockRepo.On("GetCollectionByID", ctx, 1).Return((*entities.Collection)(nil), errs.ErrCollectionNotFound).Once()

		err := usecase.DeleteCollection(ctx, 1)

		assert.ErrorIs(t, err, errs.ErrCollectionNotFound)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed to delete image", func(t *testing.T) {
		mockRepo.On("GetCollectionByID", ctx, 1).Return(&entities.Collection{
			ID:      1,
			ImageID: "image_id",
		}, nil).Once()
		mockFS.On("DeleteImage", ctx, "image_id").Return(assert.AnError).Once()

		err := usecase.DeleteCollection(ctx, 1)

		assert.Error(t, err)
		mockFS.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed to delete collection", func(t *testing.T) {
		mockRepo.On("GetCollectionByID", ctx, 1).Return(&entities.Collection{
			ID:      1,
			ImageID: "image_id",
		}, nil).Once()
		mockFS.On("DeleteImage", ctx, "image_id").Return(nil).Once()
		mockRepo.On("DeleteCollection", ctx, 1).Return(assert.AnError).Once()

		err := usecase.DeleteCollection(ctx, 1)

		assert.Error(t, err)
		mockFS.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
	})
}

type collectionsRepoMock struct {
	mock.Mock
}

func (m *collectionsRepoMock) CreateCollection(ctx context.Context, dto *dto.CreateCollectionInput) (int, error) {
	args := m.Called(ctx, dto)
	return args.Int(0), args.Error(1)
}

func (m *collectionsRepoMock) UpdateCollection(ctx context.Context, dto *dto.UpdateCollectionInput) error {
	args := m.Called(ctx, dto)
	return args.Error(0)
}

func (m *collectionsRepoMock) GetCollectionByID(ctx context.Context, id int) (*entities.Collection, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Collection), args.Error(1)
}

func (m *collectionsRepoMock) GetAllCollections(ctx context.Context) ([]*entities.Collection, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entities.Collection), args.Error(1)
}

func (m *collectionsRepoMock) DeleteCollection(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
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