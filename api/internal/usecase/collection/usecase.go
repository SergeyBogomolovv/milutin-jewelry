package collection

import (
	"context"
	"errors"
	"log/slog"
	"mime/multipart"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/collection"
)

type usecase struct {
	files   FilesService
	storage Storage
	log     *slog.Logger
}

func New(log *slog.Logger, files FilesService, storage Storage) *usecase {
	const dest = "collectionsUsecase"
	return &usecase{log: log.With(slog.String("dest", dest)), files: files, storage: storage}
}

func (u *usecase) Create(ctx context.Context, payload CreateCollectionPayload, image multipart.File) (*storage.Collection, error) {
	const op = "Create"
	log := u.log.With(slog.String("op", op))

	imageID, err := u.files.UploadImage(ctx, image, collectionsKey)
	if err != nil {
		log.Error("failed to upload image", "err", err)
		return nil, err
	}
	collection := &storage.Collection{
		Title:       payload.Title,
		Description: payload.Description,
		ImageID:     imageID,
	}
	if err := u.storage.Save(ctx, collection); err != nil {
		log.Error("failed to create collection", "err", err)
		return nil, err
	}
	return collection, nil
}

func (u *usecase) Update(ctx context.Context, payload UpdateCollectionPayload, image multipart.File) (*storage.Collection, error) {
	const op = "Update"
	log := u.log.With("op", op, "id", payload.ID)

	collection, err := u.storage.GetByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, storage.ErrCollectionNotFound) {
			return nil, ErrCollectionNotFound
		}
		log.Error("failed to get collection", "err", err)
		return nil, err
	}
	if payload.Title != nil {
		collection.Title = *payload.Title
	}
	if payload.Description != nil {
		collection.Description = *payload.Description
	}
	oldImageID := collection.ImageID
	if image != nil {
		imageID, err := u.files.UploadImage(ctx, image, collectionsKey)
		if err != nil {
			log.Error("failed to upload image", "err", err)
			return nil, err
		}
		collection.ImageID = imageID
	}
	if err := u.storage.Update(ctx, collection); err != nil {
		log.Error("failed to update collection", "err", err)
		return nil, err
	}
	if oldImageID != collection.ImageID {
		if err := u.files.DeleteImage(ctx, oldImageID); err != nil {
			log.Error("failed to delete image", "err", err)
			return nil, err
		}
	}
	return collection, nil
}

func (u *usecase) GetAll(ctx context.Context) ([]storage.Collection, error) {
	const op = "GetAll"
	log := u.log.With(slog.String("op", op))

	collections, err := u.storage.GetAll(ctx)
	if err != nil {
		log.Error("failed to get collections", "err", err)
		return nil, err
	}
	return collections, nil
}

func (u *usecase) GetByID(ctx context.Context, id int) (*storage.Collection, error) {
	const op = "GetByID"
	log := u.log.With(slog.String("op", op), slog.Int("id", id))

	collection, err := u.storage.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrCollectionNotFound) {
			return nil, ErrCollectionNotFound
		}
		log.Error("failed to get collection", "err", err)
		return nil, err
	}
	return collection, nil
}

func (u *usecase) Delete(ctx context.Context, id int) (*storage.Collection, error) {
	const op = "Delete"
	log := u.log.With(slog.String("op", op), slog.Int("id", id))

	collection, err := u.storage.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrCollectionNotFound) {
			return nil, ErrCollectionNotFound
		}
		log.Error("failed to get collection", "err", err)
		return nil, err
	}
	if err := u.storage.Delete(ctx, id); err != nil {
		log.Error("failed to delete collection", "err", err)
		return nil, err
	}
	if err := u.files.DeleteImage(ctx, collection.ImageID); err != nil {
		log.Error("failed to delete image", "err", err)
		return nil, err
	}
	return collection, nil
}
