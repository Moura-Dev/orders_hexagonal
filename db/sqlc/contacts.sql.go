// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: contacts.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createContact = `-- name: CreateContact :one
INSERT INTO contacts (company_id, user_id, email, website, address, inscricao_estadual, cnpj, name, cellphone, logo_url, fantasy_name)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id, company_id, user_id, email, website, address, inscricao_estadual, cnpj, name, cellphone, logo_url, fantasy_name, created_at, updated_at
`

type CreateContactParams struct {
	CompanyID         sql.NullInt32
	UserID            sql.NullInt32
	Email             sql.NullString
	Website           sql.NullString
	Address           sql.NullString
	InscricaoEstadual sql.NullString
	Cnpj              sql.NullString
	Name              sql.NullString
	Cellphone         sql.NullString
	LogoUrl           sql.NullString
	FantasyName       sql.NullString
}

func (q *Queries) CreateContact(ctx context.Context, arg CreateContactParams) (Contact, error) {
	row := q.db.QueryRowContext(ctx, createContact,
		arg.CompanyID,
		arg.UserID,
		arg.Email,
		arg.Website,
		arg.Address,
		arg.InscricaoEstadual,
		arg.Cnpj,
		arg.Name,
		arg.Cellphone,
		arg.LogoUrl,
		arg.FantasyName,
	)
	var i Contact
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.UserID,
		&i.Email,
		&i.Website,
		&i.Address,
		&i.InscricaoEstadual,
		&i.Cnpj,
		&i.Name,
		&i.Cellphone,
		&i.LogoUrl,
		&i.FantasyName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteContactById = `-- name: DeleteContactById :one
DELETE
FROM contacts
WHERE id = $1
    RETURNING id, company_id, user_id, email, website, address, inscricao_estadual, cnpj, name, cellphone, logo_url, fantasy_name, created_at, updated_at
`

func (q *Queries) DeleteContactById(ctx context.Context, id int32) (Contact, error) {
	row := q.db.QueryRowContext(ctx, deleteContactById, id)
	var i Contact
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.UserID,
		&i.Email,
		&i.Website,
		&i.Address,
		&i.InscricaoEstadual,
		&i.Cnpj,
		&i.Name,
		&i.Cellphone,
		&i.LogoUrl,
		&i.FantasyName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getContactCompanyId = `-- name: GetContactCompanyId :one
SELECT id, company_id, user_id, email, website, address, inscricao_estadual, cnpj, name, cellphone, logo_url, fantasy_name, created_at, updated_at
FROM contacts c
WHERE c.company_id = $1
`

func (q *Queries) GetContactCompanyId(ctx context.Context, companyID sql.NullInt32) (Contact, error) {
	row := q.db.QueryRowContext(ctx, getContactCompanyId, companyID)
	var i Contact
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.UserID,
		&i.Email,
		&i.Website,
		&i.Address,
		&i.InscricaoEstadual,
		&i.Cnpj,
		&i.Name,
		&i.Cellphone,
		&i.LogoUrl,
		&i.FantasyName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getContactsId = `-- name: GetContactsId :one
SELECT id, company_id, user_id, email, website, address, inscricao_estadual, cnpj, name, cellphone, logo_url, fantasy_name, created_at, updated_at
FROM contacts c
WHERE c.id = $1
`

func (q *Queries) GetContactsId(ctx context.Context, id int32) (Contact, error) {
	row := q.db.QueryRowContext(ctx, getContactsId, id)
	var i Contact
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.UserID,
		&i.Email,
		&i.Website,
		&i.Address,
		&i.InscricaoEstadual,
		&i.Cnpj,
		&i.Name,
		&i.Cellphone,
		&i.LogoUrl,
		&i.FantasyName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listContacts = `-- name: ListContacts :many
SELECT count(*) OVER () AS total_items, sub_query.id, sub_query.company_id, sub_query.user_id, sub_query.email, sub_query.website, sub_query.address, sub_query.inscricao_estadual, sub_query.cnpj, sub_query.name, sub_query.cellphone, sub_query.logo_url, sub_query.fantasy_name, sub_query.created_at, sub_query.updated_at
FROM (SELECT id, company_id, user_id, email, website, address, inscricao_estadual, cnpj, name, cellphone, logo_url, fantasy_name, created_at, updated_at
      FROM contacts
      ORDER BY CASE
                   WHEN NOT $3::bool AND $4::text = 'name' THEN name END,
               CASE
                   WHEN $3::bool AND $4::text = 'name' THEN name END DESC,
               CASE
                   WHEN NOT $3::bool AND $4::text = 'id' THEN id END,
               CASE
                   WHEN $3::bool AND $4::text = 'id' THEN id END DESC) sub_query
LIMIT $1
OFFSET $2
`

type ListContactsParams struct {
	Limit   int32
	Offset  int32
	Reverse bool
	OrderBy string
}

type ListContactsRow struct {
	TotalItems        int64
	ID                int32
	CompanyID         sql.NullInt32
	UserID            sql.NullInt32
	Email             sql.NullString
	Website           sql.NullString
	Address           sql.NullString
	InscricaoEstadual sql.NullString
	Cnpj              sql.NullString
	Name              sql.NullString
	Cellphone         sql.NullString
	LogoUrl           sql.NullString
	FantasyName       sql.NullString
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (q *Queries) ListContacts(ctx context.Context, arg ListContactsParams) ([]ListContactsRow, error) {
	rows, err := q.db.QueryContext(ctx, listContacts,
		arg.Limit,
		arg.Offset,
		arg.Reverse,
		arg.OrderBy,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListContactsRow{}
	for rows.Next() {
		var i ListContactsRow
		if err := rows.Scan(
			&i.TotalItems,
			&i.ID,
			&i.CompanyID,
			&i.UserID,
			&i.Email,
			&i.Website,
			&i.Address,
			&i.InscricaoEstadual,
			&i.Cnpj,
			&i.Name,
			&i.Cellphone,
			&i.LogoUrl,
			&i.FantasyName,
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

const updateContactByID = `-- name: UpdateContactByID :one
UPDATE contacts
SET company_id = COALESCE($1, company_id),
    user_id = COALESCE($2, user_id),
    email = COALESCE($3, email),
    website = COALESCE($4, website),
    address = COALESCE($5, address),
    inscricao_estadual = COALESCE($6, inscricao_estadual),
    cnpj = COALESCE($7, cnpj),
    name = COALESCE($8, name),
    cellphone = COALESCE($9, cellphone),
    logo_url = COALESCE($10, logo_url),
    fantasy_name = COALESCE($11, fantasy_name)
WHERE id = $12
RETURNING id, company_id, user_id, email, website, address, inscricao_estadual, cnpj, name, cellphone, logo_url, fantasy_name, created_at, updated_at
`

type UpdateContactByIDParams struct {
	CompanyID         sql.NullInt32
	UserID            sql.NullInt32
	Email             sql.NullString
	Website           sql.NullString
	Address           sql.NullString
	InscricaoEstadual sql.NullString
	Cnpj              sql.NullString
	Name              sql.NullString
	Cellphone         sql.NullString
	LogoUrl           sql.NullString
	FantasyName       sql.NullString
	ID                int32
}

func (q *Queries) UpdateContactByID(ctx context.Context, arg UpdateContactByIDParams) (Contact, error) {
	row := q.db.QueryRowContext(ctx, updateContactByID,
		arg.CompanyID,
		arg.UserID,
		arg.Email,
		arg.Website,
		arg.Address,
		arg.InscricaoEstadual,
		arg.Cnpj,
		arg.Name,
		arg.Cellphone,
		arg.LogoUrl,
		arg.FantasyName,
		arg.ID,
	)
	var i Contact
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.UserID,
		&i.Email,
		&i.Website,
		&i.Address,
		&i.InscricaoEstadual,
		&i.Cnpj,
		&i.Name,
		&i.Cellphone,
		&i.LogoUrl,
		&i.FantasyName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}