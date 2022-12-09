package contact

import (
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
		CompanyID:         req.CompanyID,
		UserID:            req.UserID,
		Email:             req.Email,
		Website:           req.Website,
		Address:           req.Address,
		InscricaoEstadual: req.InscricaoEstadual,
		Cnpj:              req.Cnpj,
		Name:              req.Name,
		Cellphone:         req.Cellphone,
		LogoUrl:           req.LogoUrl,
		FantasyName:       req.FantasyName,
	}

	user, err := u.db.CreateUser(ctx, arg)
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

	rsp := models.UserToJSON(user)

	ctx.JSON(http.StatusOK, rsp)
}
