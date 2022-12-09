package contact

import (
	"github.com/gin-gonic/gin"

	"project-orders/api/v1/middleware"
	"project-orders/db"
	"project-orders/token"
	"project-orders/util"
)

type IContact interface {
	SetupContactRoute(routerGroup *gin.RouterGroup)
}

type Contact struct {
	db         db.Storage
	tokenMaker token.Maker
	config     util.Config
}

func NewContact(db db.Storage, config util.Config) IContact {
	tokenMaker, _ := token.NewPasetoMaker(config.TokenSymmetricKey)

	return Contact{
		db:         db,
		config:     config,
		tokenMaker: tokenMaker,
	}
}

func (c Contact) SetupContactRoute(routerGroup *gin.RouterGroup) {
	authRoutes := routerGroup.Group("/").Use(middleware.AuthMiddleware(c.tokenMaker))
	authRoutes.POST("/contact", c.create)
	authRoutes.GET("/contact", c.list)
	authRoutes.PATCH("/contact", c.update)
	authRoutes.GET("/contact/:id", c.getByID)
	authRoutes.DELETE("/contact/:id", c.delete)
}
