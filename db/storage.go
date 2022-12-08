package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Storage interface {
	sqlc.Querier
}

type SqlStorage struct {
	db *sql.DB
	*sqlc.Queries
}

func NewStorage(db *sql.DB) Storage {
	return &SqlStorage{
		db:      db,
		Queries: sqlc.New(db),
	}
}

func (storage *SqlStorage) execTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := storage.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx errorRoute: %v, rb errorRoute: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
