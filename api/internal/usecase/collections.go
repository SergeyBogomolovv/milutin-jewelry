package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"mime/multipart"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/dto"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/entities"
	errs "github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/errors"
)

type FilesService interface {
	UploadImage(context.Context, multipart.File, string) (string, error)
	DeleteImage(context.Context, string) error
}

type CollectionsRepo interface {
	Create(context.Context, *dto.CreateCollectionInput) (int, error)
	Update(context.Context, *dto.UpdateCollectionInput) error
	Delete(context.Context, int) error
	GetByID(context.Context, int) (*entities.Collection, error)
	GetAll(context.Context) ([]*entities.Collection, error)
}

type collectionsUsecase struct {
	fs  FilesService
	cr  CollectionsRepo
	log *slog.Logger
}

func NewCollectionsUsecase(log *slog.Logger, fs FilesService, cr CollectionsRepo) *collectionsUsecase {
	return &collectionsUsecase{log: log.With(slog.String("op", "collectionsUsecase")), fs: fs, cr: cr}
}

func (u *collectionsUsecase) Create(ctx context.Context, payload *dto.CreateCollectionRequest) (int, error) {
	imageID, err := u.fs.UploadImage(ctx, payload.Image, "collections")
	if err != nil {
		u.log.Error("failed to upload image", "err", err)
		return 0, fmt.Errorf("failed to upload image: %w", err)
	}
	id, err := u.cr.Create(ctx, &dto.CreateCollectionInput{
		Title:       payload.Title,
		Description: payload.Description,
		ImageID:     imageID,
	})
	if err != nil {
		u.log.Error("failed to create collection", "err", err)
		return 0, fmt.Errorf("failed to create collection: %w", err)
	}
	return id, nil
}

func (u *collectionsUsecase) Update(ctx context.Context, payload *dto.UpdateCollectionRequest) error {
	collection, err := u.cr.GetByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, errs.ErrCollectionNotFound) {
			u.log.Info("collection not found", "err", err)
			return err
		}
		u.log.Error("failed to get collection", "err", err)
		return err
	}
	input := &dto.UpdateCollectionInput{
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
		if err := u.fs.DeleteImage(ctx, collection.ImageID); err != nil {
			u.log.Error("failed to delete image", "err", err)
			return err
		}
		imageID, err := u.fs.UploadImage(ctx, payload.Image, "collections")
		if err != nil {
			u.log.Error("failed to upload image", "err", err)
			return err
		}
		input.ImageID = imageID
	}
	if err := u.cr.Update(ctx, input); err != nil {
		u.log.Error("failed to update collection", "err", err)
		return err
	}
	return nil
}

func (u *collectionsUsecase) GetAll(ctx context.Context) ([]*dto.CollectionResponse, error) {
	collections, err := u.cr.GetAll(ctx)
	if err != nil {
		u.log.Error("failed to get collections", "err", err)
		return nil, err
	}
	res := make([]*dto.CollectionResponse, len(collections))
	for i, collection := range collections {
		res[i] = &dto.CollectionResponse{
			ID:          collection.ID,
			Title:       collection.Title,
			Description: collection.Description,
			ImageID:     collection.ImageID,
		}
	}
	return res, nil
}

func (u *collectionsUsecase) GetByID(ctx context.Context, id int) (*dto.CollectionResponse, error) {
	collection, err := u.cr.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrCollectionNotFound) {
			u.log.Info("collection not found", "err", err)
			return nil, err
		}
		u.log.Error("failed to get collection", "err", err)
		return nil, err
	}
	return &dto.CollectionResponse{
		ID:          collection.ID,
		Title:       collection.Title,
		Description: collection.Description,
		ImageID:     collection.ImageID,
	}, nil
}

func (u *collectionsUsecase) Delete(ctx context.Context, id int) error {
	collection, err := u.cr.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrCollectionNotFound) {
			u.log.Info("collection not found", "err", err)
			return err
		}
		u.log.Error("failed to get collection", "err", err)
		return err
	}
	if err := u.fs.DeleteImage(ctx, collection.ImageID); err != nil {
		u.log.Error("failed to delete image", "err", err)
		return err
	}
	if err := u.cr.Delete(ctx, id); err != nil {
		u.log.Error("failed to delete collection", "err", err)
		return err
	}
	return nil
}
