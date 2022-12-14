package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"project-orders/api/models"
	"project-orders/db/sqlc"
	"project-orders/util"
)

func (u User) create(ctx *gin.Context) {
	var req models.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	arg := sqlc.CreateUserParams{
		CompanyID: req.CompanyID,
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
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
