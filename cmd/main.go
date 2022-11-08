package main

import (
	"fmt"

	cors "github.com/gin-contrib/cors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/remove158/chula-sso/cmd/di"
)

func main() {

	server := di.InitializeEvent()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "DeeAppId", "DeeAppSecret", "DeeTicket")

	server.Gin.Use(cors.New(config))

	{
		server.AuthRoute.AddRoute()
	}

	server.Gin.LoadHTMLFiles("./templates/index.html")
	server.Gin.Run(fmt.Sprintf(":%s", server.Config.PORT))
}
