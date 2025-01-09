package banner

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var ErrBannerNotFound = errors.New("banner not found")

type storage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *storage {
	return &storage{db: db}
}

func (s *storage) Save(ctx context.Context, banner *Banner) error {
	query := `
  INSERT INTO banners (image_id, mobile_image_id, collection_id) 
  VALUES ($1, $2, $3)
  RETURNING banner_id, image_id, mobile_image_id, collection_id, created_at`
	if err := s.db.GetContext(ctx, banner, query, banner.ImageID, banner.MobileImageID, banner.CollectionID); err != nil {
		return err
	}
	return nil
}

func (s *storage) List(ctx context.Context) ([]Banner, error) {
	query := `
  SELECT banner_id, image_id, mobile_image_id, collection_id, created_at
  FROM banners 
  ORDER BY created_at DESC`
	banners := make([]Banner, 0)
	if err := s.db.SelectContext(ctx, &banners, query); err != nil {
		return nil, err
	}
	return banners, nil
}

func (s *storage) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM banners WHERE banner_id = $1`
	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if aff == 0 {
		return ErrBannerNotFound
	}
	return nil
}

func (s *storage) GetById(ctx context.Context, id int) (*Banner, error) {
	banner := new(Banner)
	query := `
	SELECT banner_id, image_id, mobile_image_id, collection_id, created_at
	FROM banners 
	WHERE banner_id = $1`
	if err := s.db.GetContext(ctx, banner, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrBannerNotFound
		}
		return nil, err
	}
	return banner, nil
}

func (r *storage) CollectionExists(ctx context.Context, id int) (bool, error) {
	var exists bool
	query := `SELECT TRUE FROM collections WHERE collection_id = $1`
	if err := r.db.GetContext(ctx, &exists, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return exists, nil
}
