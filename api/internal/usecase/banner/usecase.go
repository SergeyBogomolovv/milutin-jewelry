package banner

import (
	"context"
	"errors"
	"log/slog"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/banner"
)

var (
	ErrBannerNotFound     = errors.New("banner not found")
	ErrCollectionNotFound = errors.New("collection not found")
)

type usecase struct {
	log     *slog.Logger
	storage Storage
	files   FilesService
}

func New(log *slog.Logger, files FilesService, storage Storage) *usecase {
	const dest = "bannerUsecase"
	return &usecase{log: log.With(slog.String("dest", dest)), files: files, storage: storage}
}

func (u *usecase) List(ctx context.Context) ([]storage.Banner, error) {
	const op = "List"
	log := u.log.With(slog.String("op", op))

	banners, err := u.storage.List(ctx)
	if err != nil {
		log.Error("failed to get banner", "err", err)
		return nil, err
	}
	return banners, nil
}

func (u *usecase) Create(ctx context.Context, payload CreateBannerPayload) (*storage.Banner, error) {
	const op = "Create"
	log := u.log.With(slog.String("op", op))

	banner := new(storage.Banner)

	if payload.CollectionID != nil {
		exists, err := u.storage.CollectionExists(ctx, *payload.CollectionID)
		if err != nil {
			log.Error("can't check collection exists", "err", err)
			return nil, err
		}
		if !exists {
			return nil, ErrCollectionNotFound
		}
		banner.CollectionID = payload.CollectionID
	}

	imageID, err := u.files.UploadImage(ctx, payload.Image, bannersKey)
	if err != nil {
		log.Error("failed to upload image", "err", err)
		return nil, err
	}
	banner.ImageID = imageID

	mobileImageID, err := u.files.UploadImage(ctx, payload.MobileImage, bannersKey)
	if err != nil {
		log.Error("failed to upload image", "err", err)
		return nil, err
	}
	banner.MobileImageID = mobileImageID

	if err := u.storage.Save(ctx, banner); err != nil {
		log.Error("failed to save banner", "err", err)
		return nil, err
	}
	return banner, nil
}

func (u *usecase) Delete(ctx context.Context, id int) (*storage.Banner, error) {
	const op = "Delete"
	log := u.log.With(slog.String("op", op), slog.Int("id", id))

	banner, err := u.storage.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrBannerNotFound) {
			return nil, ErrBannerNotFound
		}
		log.Error("failed to get banner", "err", err)
		return nil, err
	}
	if err := u.files.DeleteImage(ctx, banner.ImageID); err != nil {
		log.Error("failed to delete image", "err", err, "imageID", banner.ImageID)
		return nil, err
	}
	if err := u.files.DeleteImage(ctx, banner.MobileImageID); err != nil {
		log.Error("failed to delete image", "err", err, "imageID", banner.MobileImageID)
		return nil, err
	}
	if err := u.storage.Delete(ctx, id); err != nil {
		log.Error("failed to delete banner", "err", err)
		return nil, err
	}
	return banner, nil
}
