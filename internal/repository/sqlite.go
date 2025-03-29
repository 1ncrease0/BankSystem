package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteDB(storagePath string) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
