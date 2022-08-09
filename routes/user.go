package routes

import (
	"go-api/controller"
	"github.com/gin-gonic/gin"
)


func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.PUT("/:id", controller.PutUser)
	router.DELETE("/:id", controller.DeleteUser)
}