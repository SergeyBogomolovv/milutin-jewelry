package usecase

import (
	"context"
	"log/slog"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/dto"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/entities"
)

type CollectionItemsRepo interface {
	Create(context.Context, *dto.CreateCollectionItemInput) (int, error)
	Update(context.Context, *dto.UpdateCollectionItemInput) error
	Delete(context.Context, int) error
	GetByCollection(context.Context, int) ([]*entities.CollectionItem, error)
	GetOne(context.Context, int) (*entities.CollectionItem, error)
}

type collectionItemsUsecase struct {
	repo CollectionItemsRepo
	log  *slog.Logger
	fs   FilesService
}

func NewCollectionItemsUsecase(log *slog.Logger, fs FilesService, repo CollectionItemsRepo) *collectionItemsUsecase {
	return &collectionItemsUsecase{log: log.With(slog.String("op", "collectionItemsUsecase")), repo: repo, fs: fs}
}

func (u *collectionItemsUsecase) Create(ctx context.Context, payload *dto.CreateCollectionItemRequest) (int, error) {
	return 0, nil
}

func (u *collectionItemsUsecase) Update(ctx context.Context, payload *dto.UpdateCollectionItemRequest) error {
	return nil
}

func (u *collectionItemsUsecase) Delete(ctx context.Context, id int) error {
	return nil
}

func (u *collectionItemsUsecase) GetByCollection(ctx context.Context, id int) ([]*dto.CollectionItemResponse, error) {
	return nil, nil
}

func (u *collectionItemsUsecase) GetOne(ctx context.Context, id int) (*dto.CollectionItemResponse, error) {
	return nil, nil
}
