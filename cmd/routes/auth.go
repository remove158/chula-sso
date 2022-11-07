package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/remove158/chula-sso/cmd/handlers"
)

var Set = wire.Bind(new(IAuthRoute), new(*AuthRoute))

type IAuthRoute interface {
	AddRoute()
}

type AuthRoute struct {
	authHandler handlers.IAuthHandler
	gin         *gin.Engine
}

func NewAuthRoute(authHandler handlers.IAuthHandler, gin *gin.Engine) *AuthRoute {
	return &AuthRoute{
		gin:         gin,
		authHandler: authHandler,
	}
}

func (r *AuthRoute) AddRoute() {
	r.gin.GET("/login", r.authHandler.GetLogin)
	r.gin.POST("/login", r.authHandler.PostLogin)
	r.gin.GET("/serviceValidation", r.authHandler.ServiceValidation)
	r.gin.POST("/serviceValidation", r.authHandler.ServiceValidation)
}
