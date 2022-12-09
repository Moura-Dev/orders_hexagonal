// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: companies.sql

package sqlc

import (
	"context"
)

const deleteCompany = `-- name: DeleteCompany :one
delete from companies where id = $1 returning id, name, created_at, updated_at
`

func (q *Queries) DeleteCompany(ctx context.Context, id int32) (Company, error) {
	row := q.db.QueryRowContext(ctx, deleteCompany, id)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCompanyById = `-- name: GetCompanyById :one
select id, name, created_at, updated_at from companies where id = $1
`

func (q *Queries) GetCompanyById(ctx context.Context, id int32) (Company, error) {
	row := q.db.QueryRowContext(ctx, getCompanyById, id)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertCompany = `-- name: InsertCompany :one
insert into companies (name) values ($1) returning id, name, created_at, updated_at
`

func (q *Queries) InsertCompany(ctx context.Context, name string) (Company, error) {
	row := q.db.QueryRowContext(ctx, insertCompany, name)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listCompanies = `-- name: ListCompanies :many
select id, name, created_at, updated_at from companies
`

func (q *Queries) ListCompanies(ctx context.Context) ([]Company, error) {
	rows, err := q.db.QueryContext(ctx, listCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Company{}
	for rows.Next() {
		var i Company
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
