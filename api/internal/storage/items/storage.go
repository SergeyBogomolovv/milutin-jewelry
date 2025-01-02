package itemstorage

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

func (r *storage) Save(ctx context.Context, item *Item) error {
	query := `
  INSERT INTO collection_items (collection_id, title, description, image_id) 
  VALUES ($1, $2, $3, $4) 
  RETURNING item_id, collection_id, title, description, image_id, created_at`
	if err := r.db.GetContext(ctx, item, query, item.CollectionID, item.Title, item.Description, item.ImageID); err != nil {
		return e.Wrap("can't save item", err)
	}
	return nil
}

func (r *storage) Update(ctx context.Context, item *Item) error {
	query := `
  UPDATE collection_items SET title = $1, description = $2, image_id = $3
  WHERE item_id = $4
  RETURNING item_id, collection_id, title, description, image_id, created_at`
	if err := r.db.GetContext(ctx, item, query, item.Title, item.Description, item.ImageID, item.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrItemNotFound
		}
		return e.Wrap("can't update item", err)
	}
	return nil
}

func (r *storage) Delete(ctx context.Context, id int) (err error) {
	defer func() { err = e.WrapIfErr("can't delete item", err) }()
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
		return ErrItemNotFound
	}
	return nil
}

func (r *storage) GetById(ctx context.Context, id int) (item *Item, err error) {
	defer func() { err = e.WrapIfErr("can't get item", err) }()
	item = new(Item)
	query := `
  SELECT item_id, collection_id, title, description, image_id, created_at
  FROM collection_items 
  WHERE item_id = $1`
	if err := r.db.GetContext(ctx, item, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrItemNotFound
		}
		return nil, err
	}
	return item, nil
}

func (r *storage) GetByCollectionId(ctx context.Context, id int) ([]*Item, error) {
	items := make([]*Item, 0)
	query := `
	SELECT item_id, collection_id, title, description, image_id, created_at
	FROM collection_items 
	WHERE collection_id = $1 
	ORDER BY created_at DESC`
	if err := r.db.SelectContext(ctx, &items, query, id); err != nil {
		return nil, e.Wrap("can't get items", err)
	}
	return items, nil
}

func (r *storage) CollectionExists(ctx context.Context, id int) (bool, error) {
	var exists bool
	query := `SELECT TRUE FROM collections WHERE collection_id = $1`
	if err := r.db.GetContext(ctx, &exists, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, e.Wrap("can't check collection", err)
	}
	return true, nil
}
