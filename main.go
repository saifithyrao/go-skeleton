package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rabbitmq-docker/config"
)

func main() {
	app := gin.Default()

	routes := config.NewRoutes()
	routes.SetUpRoutes(app)

	fmt.Println("start.")

	err := app.Run(":8000")
	if err != nil {
		return
	}

}
