package contact

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"project-orders/api/models"
	"project-orders/db/sqlc"
)

func (c Contact) list(ctx *gin.Context) {
	var req models.PaginationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	args := sqlc.ListContactsParams{
		Limit:   req.PageSize,
		Offset:  (req.PageID - 1) * req.PageSize,
		OrderBy: req.OrderBy,
		Reverse: req.Reverse,
	}

	contacts, err := c.db.ListContacts(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if len(contacts) == 0 {
		ctx.JSON(http.StatusOK, models.ListContactsResponse{
			Contact: []models.ContactResponse{},
			PaginationResponse: models.PaginationResponse{
				Limit:  req.PageID,
				Offset: req.PageSize,
			},
		})
		return
	}

	rsp := models.ContactsToJSONList(contacts, req.PageID, req.PageSize)

	ctx.JSON(http.StatusOK, rsp)
}
