-- name: ListCompanies :many
SELECT *
FROM companies;

-- name: GetCompanyById :one
SELECT *
FROM companies
WHERE id = $1;

-- name: CreateCompany :one
INSERT INTO companies (name)
VALUES ($1)
RETURNING *;

-- name: DeleteCompany :one
DELETE
FROM companies
WHERE id = $1
RETURNING *;