//go:build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kelseyhightower/envconfig"
	"github.com/remove158/chula-sso/cmd/handlers"
	"github.com/remove158/chula-sso/cmd/routes"
	"github.com/remove158/chula-sso/internal/services"
)

type Server struct {
	AuthRoute routes.IAuthRoute
	Gin       *gin.Engine
	Config    Setting
}

type Setting struct {
	PORT string `envconfig:"PORT" default:"8080"`
}

func InitializeServer(
	gin *gin.Engine,
	authRoute routes.IAuthRoute,
) *Server {

	var config Setting
	envconfig.Process("", &config)

	return &Server{
		Gin:       gin,
		AuthRoute: authRoute,
		Config:    config,
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
