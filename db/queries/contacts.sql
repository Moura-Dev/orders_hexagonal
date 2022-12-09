-- name: GetContactsId :one
SELECT *
FROM contacts c
WHERE c.id = $1;

-- name: GetContactCompanyId :one
SELECT *
FROM contacts c
WHERE c.company_id = $1;

-- name: ListContacts :many
SELECT count(*) OVER () AS total_items, sub_query.*
FROM (SELECT *
      FROM contacts
      ORDER BY CASE
                   WHEN NOT @reverse::bool AND @order_by::text = 'name' THEN name END,
               CASE
                   WHEN @reverse::bool AND @order_by::text = 'name' THEN name END DESC,
               CASE
                   WHEN NOT @reverse::bool AND @order_by::text = 'id' THEN id END,
               CASE
                   WHEN @reverse::bool AND @order_by::text = 'id' THEN id END DESC) sub_query


-- name: UpdateContactByID :one
UPDATE contacts
SET company_id = COALESCE(sqlc.narg('company_id'), company_id),
    user_id = COALESCE(sqlc.narg('user_id'), user_id),
    email = COALESCE(sqlc.narg('email'), email),
    website = COALESCE(sqlc.narg('website'), website),
    address = COALESCE(sqlc.narg('address'), address),
    inscricao_estadual = COALESCE(sqlc.narg('inscricao_estadual'), inscricao_estadual),
    cnpj = COALESCE(sqlc.narg('cnpj'), cnpj),
    name = COALESCE(sqlc.narg('name'), name),
    cellphone = COALESCE(sqlc.narg('cellphone'), cellphone),
    logo_url = COALESCE(sqlc.narg('logo_url'), logo_url),
    fantasy_name = COALESCE(sqlc.narg('fantasy_name'), fantasy_name)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteContactById :one
DELETE
FROM contacts
WHERE id = sqlc.arg('id')
    RETURNING *;