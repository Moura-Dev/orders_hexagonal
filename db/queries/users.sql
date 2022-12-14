-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: ListUsers :many
SELECT count(*) OVER () AS total_items, sub_query.*
FROM (SELECT *
      FROM users
      ORDER BY CASE
                   WHEN NOT @reverse::bool AND @order_by::text = 'name' THEN name END,
               CASE
                   WHEN @reverse::bool AND @order_by::text = 'name' THEN name END DESC,
               CASE
                   WHEN NOT @reverse::bool AND @order_by::text = 'id' THEN id END,
               CASE
                   WHEN @reverse::bool AND @order_by::text = 'id' THEN id END DESC) sub_query
LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (company_id, name, email, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUserByID :one
UPDATE users
SET company_id = COALESCE(sqlc.narg('company_id'), company_id),
    name       = COALESCE(sqlc.narg('name'), name),
    email      = COALESCE(sqlc.narg('email'), email),
    password   = COALESCE(sqlc.narg('password'), password)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteUserByID :one
DELETE
FROM users
WHERE id = $1
RETURNING *;