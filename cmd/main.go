package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/remove158/chula-sso/cmd/di"
)

func main() {
	server := di.InitializeEvent()

	{
		server.AuthRoute.AddRoute()
	}
	server.Gin.LoadHTMLFiles("./templates/index.html")
	server.Gin.Run(":8080")
}
