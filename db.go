package bunapp

import "github.com/uptrace/bun"

type DB struct {
	*bun.DB
}

func newDB(db *bun.DB) *DB {
	return &DB{db}
}
