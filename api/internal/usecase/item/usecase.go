package item

import (
	"context"
	"errors"
	"log/slog"
	"mime/multipart"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/item"
)

type usecase struct {
	storage Storage
	log     *slog.Logger
	files   FilesService
}

func New(log *slog.Logger, files FilesService, storage Storage) *usecase {
	const dest = "itemsUsecase"
	return &usecase{log: log.With(slog.String("dest", dest)), storage: storage, files: files}
}

func (u *usecase) Create(ctx context.Context, payload CreateItemPayload, image multipart.File) (*storage.Item, error) {
	const op = "Create"
	log := u.log.With(slog.String("op", op))

	exists, err := u.storage.CollectionExists(ctx, payload.CollectionID)
	if err != nil {
		log.Error("can't check collection exists", "err", err)
		return nil, err
	}
	if !exists {
		log.Info("collection not found")
		return nil, ErrCollectionNotFound
	}

	imageID, err := u.files.UploadImage(ctx, image, itemsKey)
	if err != nil {
		log.Error("failed to upload image", "err", err)
		return nil, err
	}

	item := &storage.Item{
		CollectionID: payload.CollectionID,
		Title:        payload.Title,
		Description:  payload.Description,
		ImageID:      imageID,
	}
	if err := u.storage.Save(ctx, item); err != nil {
		log.Error("can't save item", "err", err)
		return nil, err
	}

	return item, nil
}

func (u *usecase) Update(ctx context.Context, payload UpdateItemPayload, image multipart.File) (*storage.Item, error) {
	const op = "Update"
	log := u.log.With(slog.String("op", op))

	item, err := u.storage.GetById(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, storage.ErrItemNotFound) {
			log.Info("item not found", "err", err)
			return nil, ErrItemNotFound
		}
		log.Error("can't get item", "err", err)
		return nil, err
	}
	if payload.Title != nil {
		item.Title = *payload.Title
	}
	if payload.Description != nil {
		item.Description = *payload.Description
	}
	oldImageID := item.ImageID
	if image != nil {
		imageID, err := u.files.UploadImage(ctx, image, itemsKey)
		if err != nil {
			log.Error("can't upload image", "err", err)
			return nil, err
		}
		item.ImageID = imageID
	}
	if err := u.storage.Update(ctx, item); err != nil {
		log.Error("can't update item", "err", err)
		return nil, err
	}

	if oldImageID != item.ImageID {
		if err := u.files.DeleteImage(ctx, oldImageID); err != nil {
			log.Error("can't delete image", "err", err)
			return nil, err
		}
	}
	return item, nil
}

func (u *usecase) Delete(ctx context.Context, id int) (*storage.Item, error) {
	const op = "Delete"
	log := u.log.With(slog.String("op", op))

	item, err := u.storage.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrItemNotFound) {
			log.Info("item not found", "err", err)
			return nil, ErrItemNotFound
		}
		log.Error("can't get collection item", "err", err)
		return nil, err
	}

	if err := u.storage.Delete(ctx, id); err != nil {
		if errors.Is(err, storage.ErrItemNotFound) {
			log.Info("item not found", "err", err)
			return nil, ErrItemNotFound
		}
		log.Error("can't delete collection item", "err", err)
		return nil, err
	}
	if err := u.files.DeleteImage(ctx, item.ImageID); err != nil {
		log.Error("failed to delete image", "err", err)
		return nil, err
	}
	return item, nil
}

func (u *usecase) GetByCollectionId(ctx context.Context, id int) ([]*storage.Item, error) {
	const op = "GetByCollectionId"
	log := u.log.With(slog.String("op", op))

	exists, err := u.storage.CollectionExists(ctx, id)
	if err != nil {
		log.Error("can't check collection exists", "err", err)
		return nil, err
	}
	if !exists {
		log.Info("collection not found")
		return nil, ErrCollectionNotFound
	}

	items, err := u.storage.GetByCollectionId(ctx, id)
	if err != nil {
		log.Error("can't get items", "err", err)
		return nil, err
	}
	return items, nil
}
