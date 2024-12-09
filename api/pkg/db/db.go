package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(url string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}
