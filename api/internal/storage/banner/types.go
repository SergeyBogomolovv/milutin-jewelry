package banner

import "time"

type Banner struct {
	ID            int       `db:"banner_id"`
	ImageID       string    `db:"image_id"`
	MobileImageID string    `db:"mobile_image_id"`
	CollectionID  *int      `db:"collection_id"`
	CreatedAt     time.Time `db:"created_at"`
}
