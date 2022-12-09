-- name: ListCompanies :many
select * from companies;


-- name: GetCompanyById :one
select * from companies where id = $1;

-- name: InsertCompany :one
insert into companies (name) values ($1) returning *;

-- name: DeleteCompany :one
delete from companies where id = $1 returning *;