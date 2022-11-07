//go:build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/remove158/chula-sso/cmd/handlers"
	"github.com/remove158/chula-sso/cmd/routes"
	"github.com/remove158/chula-sso/internal/services"
)

type Server struct {
	AuthRoute routes.IAuthRoute
	Gin       *gin.Engine
}

func InitializeServer(
	gin *gin.Engine,
	authRoute routes.IAuthRoute,
) *Server {
	return &Server{
		Gin:       gin,
		AuthRoute: authRoute,
	}
}

func InitializeGin() *gin.Engine {
	return gin.Default()
}

func InitializeEvent() *Server {
	wire.Build(
		handlers.Set, services.Set, routes.Set,
		handlers.NewAuthHandler, services.NewAuthService, routes.NewAuthRoute,
		InitializeGin, InitializeServer,
	)
	return &Server{}
}
