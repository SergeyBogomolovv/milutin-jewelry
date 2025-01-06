package items_test

import (
	"context"
	"mime/multipart"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/item"
	"github.com/stretchr/testify/mock"
)

type storageMock struct {
	mock.Mock
}

func (m *storageMock) Save(ctx context.Context, item *storage.Item) error {
	args := m.Called(ctx, item)
	return args.Error(0)
}

func (m *storageMock) Update(ctx context.Context, item *storage.Item) error {
	args := m.Called(ctx, item)
	return args.Error(0)
}

func (m *storageMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *storageMock) GetById(ctx context.Context, id int) (*storage.Item, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*storage.Item), args.Error(1)
}

func (m *storageMock) GetByCollectionId(ctx context.Context, id int) ([]storage.Item, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]storage.Item), args.Error(1)
}

func (m *storageMock) CollectionExists(ctx context.Context, id int) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}

type filesServiceMock struct {
	mock.Mock
}

func (m *filesServiceMock) DeleteImage(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

func (m *filesServiceMock) UploadImage(ctx context.Context, file multipart.File, dest string) (string, error) {
	args := m.Called(ctx, file, dest)
	return args.String(0), args.Error(1)
}
