package banner

import (
	"context"

	storage "github.com/SergeyBogomolovv/milutin-jewelry/internal/storage/banner"
	uc "github.com/SergeyBogomolovv/milutin-jewelry/internal/usecase/banner"
)

type Usecase interface {
	List(ctx context.Context) ([]storage.Banner, error)
	Create(ctx context.Context, payload uc.CreateBannerPayload) (*storage.Banner, error)
	Delete(ctx context.Context, id int) (*storage.Banner, error)
}

type Banner struct {
	ID            int    `json:"id"`
	ImageID       string `json:"image_id"`
	MobileImageID string `json:"mobile_image_id"`
	CollectionID  int    `json:"collection_id,omitempty"`
}

func NewBanner(banner storage.Banner) Banner {
	res := Banner{
		ID:            banner.ID,
		ImageID:       banner.ImageID,
		MobileImageID: banner.MobileImageID,
	}
	if banner.CollectionID != nil {
		res.CollectionID = *banner.CollectionID
	}
	return res
}
