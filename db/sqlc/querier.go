// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	CreateCompany(ctx context.Context, name string) (Company, error)
	CreateContact(ctx context.Context, arg CreateContactParams) (Contact, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCompany(ctx context.Context, id int32) (Company, error)
	DeleteContactById(ctx context.Context, id int32) (Contact, error)
	DeleteUserByID(ctx context.Context, id int32) (User, error)
	GetCompanyById(ctx context.Context, id int32) (Company, error)
	GetContactByCompanyID(ctx context.Context, companyID sql.NullInt32) (Contact, error)
	GetContactByID(ctx context.Context, id int32) (Contact, error)
	GetSessionByID(ctx context.Context, id uuid.UUID) (Session, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int32) (User, error)
	ListCompanies(ctx context.Context) ([]Company, error)
	ListContacts(ctx context.Context, arg ListContactsParams) ([]ListContactsRow, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]ListUsersRow, error)
	UpdateContactByID(ctx context.Context, arg UpdateContactByIDParams) (Contact, error)
	UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) (User, error)
}

var _ Querier = (*Queries)(nil)
