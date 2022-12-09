package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"project-orders/api/v1/contact"
	"project-orders/api/v1/paseto"
	"project-orders/api/v1/user"
	"project-orders/db"
	"project-orders/token"
	"project-orders/util"
)

type Server struct {
	storage    db.Storage
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

func NewGin(config util.Config, storage db.Storage) *Server {
	tokenMaker, _ := token.NewPasetoMaker(config.TokenSymmetricKey)

	server := &Server{
		config:     config,
		storage:    storage,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	server.createRoutesV1(router)

	server.router = router
}

func (server *Server) createRoutesV1(router *gin.Engine) {
	router.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	v1 := router.Group("/v1")

	userRoutes := user.NewUser(server.storage, server.config)
	pasetoRoutes := paseto.NewPaseto(server.storage)
	contactRoutes := contact.NewContact(server.storage, server.config)

	userRoutes.SetupUserRoute(v1)
	pasetoRoutes.SetupPasetoRoute(v1)
	contactRoutes.SetupContactRoute(v1)
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
