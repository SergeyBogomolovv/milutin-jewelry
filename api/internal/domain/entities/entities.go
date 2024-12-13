package entities

type Collection struct {
	ID          int    `db:"collection_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	ImageID     string `db:"image_id"`
}

type CollectionItem struct {
	ID           int    `db:"collection_item_id"`
	CollectionID int    `db:"collection_id"`
	Title        string `db:"title"`
	Description  string `db:"description"`
	ImageID      string `db:"image_id"`
}
