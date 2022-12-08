package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orders_hexagonal/db"
	"orders_hexagonal/token"
	"orders_hexagonal/util"
)

type Server struct {
	storage    db.Storage
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, storage db.Storage) *Server {
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

	tokenRoutes := paseto.NewToken(server.storage)
	userRoutes.SetupUserRoute(v1)
	tokenRoutes.SetupTokenRoute(v1)

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
