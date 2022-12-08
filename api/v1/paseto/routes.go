package paseto

import (
	"github.com/gin-gonic/gin"

	"project-orders/db"
	"project-orders/token"
	"project-orders/util"
)

type IPaseto interface {
	SetupPasetoRoute(routerGroup *gin.RouterGroup)
}

type Paseto struct {
	db         db.Storage
	tokenMaker token.Maker
	config     util.Config
}

func NewPaseto(db db.Storage) IPaseto {
	return Paseto{
		db: db,
	}
}

func (p Paseto) SetupPasetoRoute(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/token", p.renewAccessToken)
}
