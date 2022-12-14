package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"project-orders/api/models"
)

func (u User) getByID(ctx *gin.Context) {
	var req models.GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := u.db.GetUserByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}

		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := models.UserToJSON(user)

	ctx.JSON(http.StatusOK, rsp)
}
