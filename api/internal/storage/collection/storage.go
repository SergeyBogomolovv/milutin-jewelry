package collection

import (
	"context"
	"database/sql"
	"errors"

	"github.com/SergeyBogomolovv/milutin-jewelry/pkg/lib/e"
	"github.com/jmoiron/sqlx"
)

type storage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *storage {
	return &storage{db: db}
}

func (r *storage) Save(ctx context.Context, collection *Collection) error {
	query := `
  INSERT INTO collections (title, description, image_id) 
  VALUES ($1, $2, $3) 
  RETURNING collection_id, title, description, image_id, created_at`
	if err := r.db.GetContext(ctx, collection, query, collection.Title, collection.Description, collection.ImageID); err != nil {
		return e.Wrap("can't save collection", err)
	}
	return nil
}

func (r *storage) Update(ctx context.Context, collection *Collection) (err error) {
	defer func() { err = e.WrapIfErr("can't update collection", err) }()
	query := `
  UPDATE collections SET title = $1, description = $2, image_id = $3 
  WHERE collection_id = $4
  RETURNING collection_id, title, description, image_id, created_at`
	if err := r.db.GetContext(ctx, collection, query, collection.Title, collection.Description, collection.ImageID, collection.ID); err != nil {
		return err
	}
	return nil
}

func (r *storage) GetByID(ctx context.Context, id int) (*Collection, error) {
	collection := new(Collection)
	query := `
  SELECT collection_id, title, description, image_id, created_at
  FROM collections 
  WHERE collection_id = $1`
	if err := r.db.GetContext(ctx, collection, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrCollectionNotFound
		}
		return nil, err
	}
	return collection, nil
}

func (r *storage) GetAll(ctx context.Context) ([]*Collection, error) {
	collections := make([]*Collection, 0)
	query := `
  SELECT collection_id, title, description, image_id, created_at
  FROM collections 
  ORDER BY created_at DESC`
	if err := r.db.SelectContext(ctx, &collections, query); err != nil {
		return nil, e.Wrap("can't get collections", err)
	}
	return collections, nil
}

func (r *storage) Delete(ctx context.Context, id int) (err error) {
	defer func() { err = e.WrapIfErr("can't delete collection", err) }()
	query := `DELETE FROM collections WHERE collection_id = $1`
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if aff == 0 {
		return ErrCollectionNotFound
	}
	return nil
}
