package usecase

import (
	"context"
	"errors"
	"log/slog"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/dto"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/entities"
	errs "github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/errors"
)

type CollectionItemsRepo interface {
	Create(context.Context, *dto.CreateCollectionItemInput) (int, error)
	Update(context.Context, *dto.UpdateCollectionItemInput) error
	Delete(context.Context, int) error
	GetByCollection(context.Context, int) ([]*entities.CollectionItem, error)
	GetOne(context.Context, int) (*entities.CollectionItem, error)
	CollectionExists(context.Context, int) error
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
	if err := u.repo.CollectionExists(ctx, payload.CollectionID); err != nil {
		if errors.Is(err, errs.ErrCollectionNotFound) {
			u.log.Info("collection not found", "err", err)
			return 0, err
		}
		u.log.Error("failed to check collection exists", "err", err)
		return 0, err
	}

	imageID, err := u.fs.UploadImage(ctx, payload.Image, "collection-items")
	if err != nil {
		u.log.Error("failed to upload image", "err", err)
		return 0, err
	}

	id, err := u.repo.Create(ctx, &dto.CreateCollectionItemInput{
		CollectionID: payload.CollectionID,
		Title:        payload.Title,
		Description:  payload.Description,
		ImageID:      imageID,
	})
	if err != nil {
		if err := u.fs.DeleteImage(ctx, imageID); err != nil {
			u.log.Error("failed to delete image", "err", err)
		}
		u.log.Error("failed to create collection item", "err", err)
		return 0, err
	}
	return id, nil
}

func (u *collectionItemsUsecase) Update(ctx context.Context, payload *dto.UpdateCollectionItemRequest) error {
	collection, err := u.repo.GetOne(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, errs.ErrCollectionItemNotFound) {
			u.log.Info("collection item not found", "err", err)
			return err
		}
		u.log.Error("failed to get collection item", "err", err)
		return err
	}

	input := &dto.UpdateCollectionItemInput{
		Title:       collection.Title,
		Description: collection.Description,
		ImageID:     collection.ImageID,
		ID:          collection.ID,
	}
	if payload.Title != "" {
		input.Title = payload.Title
	}
	if payload.Description != "" {
		input.Description = payload.Description
	}
	if payload.Image != nil {
		imageID, err := u.fs.UploadImage(ctx, payload.Image, "collection-items")
		if err != nil {
			u.log.Error("failed to upload image", "err", err)
			return err
		}
		input.ImageID = imageID
	}
	if err := u.repo.Update(ctx, input); err != nil {
		u.log.Error("failed to update collection item", "err", err)
		return err
	}
	if payload.Image != nil {
		if err := u.fs.DeleteImage(ctx, collection.ImageID); err != nil {
			u.log.Error("failed to delete image", "err", err)
		}
	}
	return nil
}

func (u *collectionItemsUsecase) Delete(ctx context.Context, id int) error {
	collection, err := u.repo.GetOne(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrCollectionItemNotFound) {
			u.log.Info("collection item not found", "err", err)
			return err
		}
		u.log.Error("failed to get collection item", "err", err)
		return err
	}

	if err := u.repo.Delete(ctx, id); err != nil {
		u.log.Error("failed to delete collection item", "err", err)
		return err
	}
	if err := u.fs.DeleteImage(ctx, collection.ImageID); err != nil {
		u.log.Error("failed to delete image", "err", err)
	}
	return nil
}

func (u *collectionItemsUsecase) GetByCollection(ctx context.Context, id int) ([]*dto.CollectionItemResponse, error) {
	if err := u.repo.CollectionExists(ctx, id); err != nil {
		if errors.Is(err, errs.ErrCollectionNotFound) {
			u.log.Info("collection not found", "err", err)
			return nil, err
		}
		u.log.Error("failed to check collection exists", "err", err)
		return nil, err
	}
	collections, err := u.repo.GetByCollection(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrCollectionNotFound) {
			u.log.Info("collection not found", "err", err)
			return nil, err
		}
		u.log.Error("failed to get collection", "err", err)
		return nil, err
	}

	var res []*dto.CollectionItemResponse
	for _, collection := range collections {
		res = append(res, &dto.CollectionItemResponse{
			ID:          collection.ID,
			Title:       collection.Title,
			Description: collection.Description,
			ImageID:     collection.ImageID,
		})
	}
	return res, nil
}

func (u *collectionItemsUsecase) GetOne(ctx context.Context, id int) (*dto.CollectionItemResponse, error) {
	collection, err := u.repo.GetOne(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrCollectionItemNotFound) {
			u.log.Info("collection item not found", "err", err)
			return nil, err
		}
		u.log.Error("failed to get collection item", "err", err)
		return nil, err
	}
	return &dto.CollectionItemResponse{
		ID:          collection.ID,
		Title:       collection.Title,
		Description: collection.Description,
		ImageID:     collection.ImageID,
	}, nil
}
