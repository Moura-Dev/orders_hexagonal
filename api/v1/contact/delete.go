package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"project-orders/api/models"
)

func (u Contact) delete(ctx *gin.Context) {
	var req models.DeleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	_, err := u.db.DeleteUserByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "success")
}
