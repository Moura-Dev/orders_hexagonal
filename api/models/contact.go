package models

import (
	"math"

	"project-orders/db/sqlc"
)

type CreateContactRequest struct {
	CompanyID         int32  `json:"companyId" binding:"required"`
	UserID            int32  `json:"userId" binding:"required"`
	Email             string `json:"email" binding:"required"`
	Website           string `json:"website" binding:"required"`
	Address           string `json:"address" binding:"required"`
	InscricaoEstadual string `json:"inscricaoEstadual" binding:"required"`
	CNPJ              string `json:"cnpj" binding:"required"`
	Name              string `json:"name" binding:"required"`
	Cellphone         string `json:"cellphone" binding:"required"`
	LogoURL           string `json:"logoUrl" binding:"required"`
	FantasyName       string `json:"fantasyName" binding:"required"`
}

type DeleteContactRequest struct {
	ID int32 `uri:"id" uri:"id" binding:"required,min=1"`
}

type GetContactRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type ContactResponse struct {
	ID                int32  `json:"id"`
	CompanyID         int32  `json:"company_id"`
	UserID            int32  `json:"userId"`
	Email             string `json:"email"`
	Website           string `json:"website"`
	Address           string `json:"address"`
	InscricaoEstadual string `json:"inscricaoEstadual"`
	CNPJ              string `json:"cnpj"`
	Name              string `json:"name"`
	Cellphone         string `json:"cellphone"`
	LogoURL           string `json:"logo_url"`
	FantasyName       string `json:"fantasy_name"`
	CreatedAt         int64  `json:"created_at"`
	UpdatedAt         int64  `json:"updated_at"`
}

type ListContactsResponse struct {
	Contact []ContactResponse `json:"users"`
	PaginationResponse
}

func ContactToJSON(SQLContact sqlc.Contact) ContactResponse {
	return ContactResponse{
		ID:          SQLContact.ID,
		CompanyID:   SQLContact.CompanyID.Int32,
		UserID:      SQLContact.UserID.Int32,
		Email:       SQLContact.Email.String,
		Website:     SQLContact.Website.String,
		Address:     SQLContact.Address.String,
		CNPJ:        SQLContact.Cnpj.String,
		Name:        SQLContact.Name.String,
		Cellphone:   SQLContact.Cellphone.String,
		LogoURL:     SQLContact.LogoUrl.String,
		FantasyName: SQLContact.FantasyName.String,
		CreatedAt:   SQLContact.CreatedAt.Unix(),
		UpdatedAt:   SQLContact.UpdatedAt.Unix(),
	}
}

func ContactsToJSONList(SQLContacts []sqlc.ListContactsRow, pageID, pageSize int32) ListContactsResponse {
	var contatcs []ContactResponse

	for _, contact := range SQLContacts {
		contatcs = append(contatcs, ContactResponse{
			ID:                contact.ID,
			CompanyID:         contact.CompanyID.Int32,
			UserID:            contact.UserID.Int32,
			Email:             contact.Email.String,
			Website:           contact.Website.String,
			Address:           contact.Address.String,
			InscricaoEstadual: contact.InscricaoEstadual.String,
			CNPJ:              contact.Cnpj.String,
			Name:              contact.Name.String,
			Cellphone:         contact.Cellphone.String,
			LogoURL:           contact.LogoUrl.String,
			FantasyName:       contact.FantasyName.String,
			CreatedAt:         contact.CreatedAt.Unix(),
			UpdatedAt:         contact.UpdatedAt.Unix(),
		})
	}

	totalPages := int32(math.Ceil(float64(SQLContacts[0].TotalItems) / float64(pageSize)))

	return ListContactsResponse{
		Contact: contatcs,
		PaginationResponse: PaginationResponse{
			Limit:      pageID,
			Offset:     pageSize,
			TotalItems: SQLContacts[0].TotalItems,
			TotalPages: totalPages,
		},
	}
}
