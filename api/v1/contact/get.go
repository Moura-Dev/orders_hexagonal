package contact

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"project-orders/api/models"
)

func (c Contact) getByID(ctx *gin.Context) {
	var req models.GetContactRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	contact, err := c.db.GetContactByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}

		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := models.ContactToJSON(contact)

	ctx.JSON(http.StatusOK, rsp)
}
