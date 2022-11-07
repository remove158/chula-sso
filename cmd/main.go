package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/remove158/chula-sso/cmd/di"
)

func main() {

	server := di.InitializeEvent()

	{
		server.AuthRoute.AddRoute()
	}

	server.Gin.LoadHTMLFiles("./templates/index.html")
	server.Gin.Run(fmt.Sprintf(":%s", server.Config.PORT))
}
