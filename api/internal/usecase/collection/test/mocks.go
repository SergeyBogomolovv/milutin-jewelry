package collection_test

import (
	"context"
	"mime/multipart"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/collection"
	"github.com/stretchr/testify/mock"
)

type storageMock struct {
	mock.Mock
}

func (m *storageMock) Save(ctx context.Context, collection *storage.Collection) error {
	args := m.Called(ctx, collection)
	return args.Error(0)
}

func (m *storageMock) Update(ctx context.Context, collection *storage.Collection) error {
	args := m.Called(ctx, collection)
	return args.Error(0)
}

func (m *storageMock) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *storageMock) GetByID(ctx context.Context, id int) (*storage.Collection, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*storage.Collection), args.Error(1)
}

func (m *storageMock) GetAll(ctx context.Context) ([]*storage.Collection, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*storage.Collection), args.Error(1)
}

type fileServiceMock struct {
	mock.Mock
}

func (m *fileServiceMock) DeleteImage(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

func (m *fileServiceMock) UploadImage(ctx context.Context, file multipart.File, key string) (string, error) {
	args := m.Called(ctx, file, key)
	return args.String(0), args.Error(1)
}
