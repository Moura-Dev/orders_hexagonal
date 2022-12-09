package contact

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"project-orders/api/models"
	"project-orders/db/sqlc"
)

func (c Contact) create(ctx *gin.Context) {
	var req models.CreateContactRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := sqlc.CreateContactParams{
		CompanyID:         sql.NullInt32{Int32: req.CompanyID, Valid: true},
		UserID:            sql.NullInt32{Int32: req.UserID, Valid: true},
		Email:             sql.NullString{String: req.Email, Valid: true},
		Website:           sql.NullString{String: req.Website, Valid: true},
		Address:           sql.NullString{String: req.Address, Valid: true},
		InscricaoEstadual: sql.NullString{String: req.InscricaoEstadual, Valid: true},
		Cnpj:              sql.NullString{String: req.CNPJ, Valid: true},
		Name:              sql.NullString{String: req.Name, Valid: true},
		Cellphone:         sql.NullString{String: req.Cellphone, Valid: true},
		LogoUrl:           sql.NullString{String: req.LogoURL, Valid: true},
		FantasyName:       sql.NullString{String: req.FantasyName, Valid: true},
	}

	contact, err := c.db.CreateContact(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, err.Error())
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := models.ContactToJSON(contact)

	ctx.JSON(http.StatusOK, rsp)
}
