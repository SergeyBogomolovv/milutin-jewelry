package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(url string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %s", err)
	}

	return db
}
