package repo

import (
	"context"

	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/dto"
	"github.com/SergeyBogomolovv/milutin-jewelry/internal/domain/entities"
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

func (r *collectionsRepo) UpdateCollection(ctx context.Context, payload *dto.UpdateCollectionInput, id int) error {
	query := `UPDATE collections SET title = $1, description = $2, image_id = $3 WHERE id = $4`
	if _, err := r.db.ExecContext(ctx, query, payload.Title, payload.Description, payload.ImageID, id); err != nil {
		return err
	}
	return nil
}

func (r *collectionsRepo) GetAllCollections(ctx context.Context) ([]*entities.Collection, error) {
	return nil, nil
}
