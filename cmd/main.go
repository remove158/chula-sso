package main

import (
	"fmt"

	cors "github.com/gin-contrib/cors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/remove158/chula-sso/cmd/di"
	healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	health_config "github.com/tavsec/gin-healthcheck/config"
)

func main() {

	server := di.InitializeEvent()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "DeeAppId", "DeeAppSecret", "DeeTicket")
	healthcheck.New(server.Gin, health_config.DefaultConfig(), []checks.Check{})

	server.Gin.Use(cors.New(config))

	{
		server.AuthRoute.AddRoute()
	}

	server.Gin.LoadHTMLFiles("./templates/index.html")
	server.Gin.Run(fmt.Sprintf(":%s", server.Config.PORT))
}
