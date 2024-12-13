package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/dto"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/entities"
	errs "github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/errors"
	"github.com/jmoiron/sqlx"
)

type collectionsRepo struct {
	db *sqlx.DB
}

func NewCollectionsRepo(db *sqlx.DB) *collectionsRepo {
	return &collectionsRepo{db: db}
}

func (r *collectionsRepo) CreateCollection(ctx context.Context, payload *dto.CreateCollectionInput) (int, error) {
	var id int
	query := `INSERT INTO collections (title, description, image_id) VALUES ($1, $2, $3) RETURNING collection_id`
	if err := r.db.GetContext(ctx, &id, query, payload.Title, payload.Description, payload.ImageID); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *collectionsRepo) UpdateCollection(ctx context.Context, payload *dto.UpdateCollectionInput) error {
	query := `UPDATE collections SET title = $1, description = $2, image_id = $3 WHERE collection_id = $4`
	if _, err := r.db.ExecContext(ctx, query, payload.Title, payload.Description, payload.ImageID, payload.ID); err != nil {
		return err
	}
	return nil
}

func (r *collectionsRepo) GetCollectionByID(ctx context.Context, id int) (*entities.Collection, error) {
	collection := new(entities.Collection)
	query := `SELECT collection_id, title, description, image_id FROM collections WHERE collection_id = $1`
	if err := r.db.GetContext(ctx, collection, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrCollectionNotFound
		}
		return nil, err
	}
	return collection, nil
}

func (r *collectionsRepo) GetAllCollections(ctx context.Context) ([]*entities.Collection, error) {
	collections := make([]*entities.Collection, 0)
	query := `SELECT collection_id, title, description, image_id FROM collections ORDER BY collection_id`
	if err := r.db.SelectContext(ctx, &collections, query); err != nil {
		return nil, err
	}
	return collections, nil
}

func (r *collectionsRepo) DeleteCollection(ctx context.Context, id int) error {
	query := `DELETE FROM collections WHERE collection_id = $1`
	if res, err := r.db.ExecContext(ctx, query, id); err != nil {
		if aff, err := res.RowsAffected(); err == nil && aff == 0 {
			return errs.ErrCollectionNotFound
		}
		return err
	}
	return nil
}
