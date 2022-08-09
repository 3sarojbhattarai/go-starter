package main

import (
	"go-api/config"
	"go-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	routes.UserRoute(router)
	config.Connect()
	router.Run(":8080")
}
