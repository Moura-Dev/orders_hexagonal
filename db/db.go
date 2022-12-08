package db

import (
	"context"
	"database/sql"
	"fmt"

	"project-orders/db/sqlc"
)

type Storage interface {
	sqlc.Querier
}

type SQLStore struct {
	db *sql.DB
	*sqlc.Queries
}

func NewStore(db *sql.DB) Storage {
	return &SQLStore{
		db:      db,
		Queries: sqlc.New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
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
