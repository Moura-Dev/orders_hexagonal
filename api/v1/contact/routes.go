package user

import (
	"github.com/gin-gonic/gin"

	"project-orders/api/v1/middleware"
	"project-orders/db"
	"project-orders/token"
	"project-orders/util"
)

type IUser interface {
	SetupUserRoute(routerGroup *gin.RouterGroup)
}

type Contact struct {
	db         db.Storage
	tokenMaker token.Maker
	config     util.Config
}

func NewUser(db db.Storage, config util.Config) IUser {
	tokenMaker, _ := token.NewPasetoMaker(config.TokenSymmetricKey)

	return Contact{
		db:         db,
		config:     config,
		tokenMaker: tokenMaker,
	}
}

func (u Contact) SetupUserRoute(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/user", c.create)

	authRoutes := routerGroup.Group("/").Use(middleware.AuthMiddleware(u.tokenMaker))
	authRoutes.GET("/user", u.list)
	authRoutes.PATCH("/user", u.update)
	authRoutes.GET("/user/:id", u.getByID)
	authRoutes.DELETE("/user/:id", u.delete)
}
