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

type collectionItemsRepo struct {
	db *sqlx.DB
}

func NewCollectionItemsRepo(db *sqlx.DB) *collectionItemsRepo {
	return &collectionItemsRepo{db: db}
}

func (r *collectionItemsRepo) Create(ctx context.Context, payload *dto.CreateCollectionItemInput) (int, error) {
	var id int
	query := `INSERT INTO collection_items (collection_id, title, description, image_id) VALUES ($1, $2, $3, $4) RETURNING item_id`
	if err := r.db.GetContext(ctx, &id, query, payload.CollectionID, payload.Title, payload.Description, payload.ImageID); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *collectionItemsRepo) Update(ctx context.Context, payload *dto.UpdateCollectionItemInput) error {
	query := `UPDATE collection_items SET title = $1, description = $2, image_id = $3 WHERE item_id = $4`
	res, err := r.db.ExecContext(ctx, query, payload.Title, payload.Description, payload.ImageID, payload.ID)
	if err != nil {
		return err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if aff == 0 {
		return errs.ErrCollectionItemNotFound
	}
	return nil
}

func (r *collectionItemsRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM collection_items WHERE item_id = $1`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if aff == 0 {
		return errs.ErrCollectionItemNotFound
	}
	return nil
}

func (r *collectionItemsRepo) GetOne(ctx context.Context, id int) (*entities.CollectionItem, error) {
	collection := new(entities.CollectionItem)
	query := `SELECT item_id, collection_id, title, description, image_id FROM collection_items WHERE item_id = $1`
	if err := r.db.GetContext(ctx, collection, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrCollectionItemNotFound
		}
		return nil, err
	}
	return collection, nil
}

func (r *collectionItemsRepo) GetByCollection(ctx context.Context, id int) ([]*entities.CollectionItem, error) {
	collections := make([]*entities.CollectionItem, 0)
	query := `SELECT item_id, collection_id, title, description, image_id FROM collection_items WHERE collection_id = $1`
	if err := r.db.SelectContext(ctx, &collections, query, id); err != nil {
		return nil, err
	}
	return collections, nil
}

func (r *collectionItemsRepo) CollectionExists(ctx context.Context, id int) error {
	var exists bool
	query := `SELECT TRUE FROM collections WHERE collection_id = $1`
	if err := r.db.GetContext(ctx, &exists, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrCollectionNotFound
		}
		return err
	}
	return nil
}
