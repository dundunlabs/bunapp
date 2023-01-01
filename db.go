package bunapp

import (
	"context"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

type DB struct {
	*bun.DB
}

func newDB(db *bun.DB) *DB {
	return &DB{db}
}

func (db *DB) AutoMigrate(ctx context.Context, migrations *migrate.Migrations) error {
	migrator := migrate.NewMigrator(db.DB, migrations)

	if err := migrator.Init(ctx); err != nil {
		return err
	}

	if err := migrator.Lock(ctx); err != nil {
		return err
	}
	defer migrator.Unlock(ctx)

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}
	if group.IsZero() {
		log.Println("there are no new migrations to run (database is up to date)")
		return nil
	}
	log.Println("migrated to", group)
	return nil
}
