package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"mime/multipart"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/dto"
)

type FilesService interface {
	UploadImage(context.Context, multipart.File, string) (string, error)
}

type CollectionsRepo interface {
	CreateCollection(context.Context, *dto.CreateCollectionInput) (int, error)
	UpdateCollection(context.Context, *dto.UpdateCollectionInput, int) error
}

type collectionsUsecase struct {
	fs  FilesService
	cr  CollectionsRepo
	log *slog.Logger
}

func NewCollectionsUsecase(log *slog.Logger, fs FilesService, cr CollectionsRepo) *collectionsUsecase {
	return &collectionsUsecase{log: log.With(slog.String("op", "collectionsUsecase")), fs: fs, cr: cr}
}

func (u *collectionsUsecase) CreateCollection(ctx context.Context, payload *dto.CreateCollectionRequest) (int, error) {
	imageID, err := u.fs.UploadImage(ctx, payload.Image, "collections")
	if err != nil {
		u.log.Error("failed to upload image", "err", err)
		return 0, fmt.Errorf("failed to upload image: %w", err)
	}
	id, err := u.cr.CreateCollection(ctx, &dto.CreateCollectionInput{
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

func (u *collectionsUsecase) UpdateCollection(ctx context.Context, payload *dto.UpdateCollectionRequest, id int) error {
	input := &dto.UpdateCollectionInput{Title: payload.Title, Description: payload.Description}
	if payload.ImageHeader != nil {
		//TODO: delete old image
		// imageID, err := u.fs.UploadImage(ctx, payload.ImageHeader, "collections")
		// if err != nil {
		// 	u.log.Error("failed to upload image", "err", err)
		// 	return err
		// }
		// input.ImageID = imageID
	}
	if err := u.cr.UpdateCollection(ctx, input, id); err != nil {
		u.log.Error("failed to update collection", "err", err)
		return err
	}
	return nil
}

func (u *collectionsUsecase) GetAllCollections(ctx context.Context) ([]*dto.CollectionResponse, error) {
	return nil, nil
}
