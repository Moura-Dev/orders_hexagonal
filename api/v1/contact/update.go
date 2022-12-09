package contact

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"project-orders/api/models"
	"project-orders/db/sqlc"
)

func (c Contact) update(ctx *gin.Context) {
	var req models.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := sqlc.UpdateUserByIDParams{
		ID:       req.ID,
		Name:     sql.NullString{String: req.Name, Valid: req.Name != ""},
		Email:    sql.NullString{String: req.Email, Valid: req.Email != ""},
		Password: sql.NullString{String: req.Password, Valid: req.Password != ""},
	}

	contact, err := c.db.UpdateUserByID(ctx, arg)
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

	rsp := models.UserToJSON(contact)

	ctx.JSON(http.StatusOK, rsp)
}
