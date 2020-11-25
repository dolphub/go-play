package database

import (
	"errors"

	"github.com/go-pg/pg/v10"
)

func New() (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		User:     "user",
		Password: "pass",
		Database: "core",
	})

	res, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	if res.RowsReturned() != 1 {
		return nil, errors.New("Database not connected")
	}
	return db, nil
}
