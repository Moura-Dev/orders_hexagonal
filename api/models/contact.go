package models

type CreateContactRequest struct {
	CompanyID         int64  `json:"company_id" binding:"required"`
	UserID            int64  `json:"user_id" binding:"required"`
	Email             string `json:"email" binding:"required"`
	Website           string `json:"website" binding:"required"`
	Address           string `json:"address" binding:"required"`
	InscricaoEstadual string `json:"inscricao_estadual" binding:"required"`
	CNPJ              string `json:"cnpj" binding:"required"`
	Name              string `json:"name" binding:"required"`
	Cellphone         string `json:"cellphone" binding:"required"`
	LogoURL           string `json:"logo_url" binding:"required"`
	FantasyName       string `json:"fantasy_name" binding:"required"`
}

type ResponseContact struct {
	ID                int64  `json:"id"`
	UserID            int64  `json:"user_id" binding:"required"`
	Email             string `json:"email" binding:"required"`
	Website           string `json:"website" binding:"required"`
	Address           string `json:"address" binding:"required"`
	InscricaoEstadual string `json:"inscricao_estadual" binding:"required"`
	CNPJ              string `json:"cnpj" binding:"required"`
	Name              string `json:"name" binding:"required"`
	Cellphone         string `json:"cellphone" binding:"required"`
	LogoURL           string `json:"logo_url" binding:"required"`
	FantasyName       string `json:"fantasy_name" binding:"required"`
}

}
