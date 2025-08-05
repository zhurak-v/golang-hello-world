package routes

import (
	"golang-hello-world/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	var router *gin.Engine = gin.Default()
	router.GET("/hello", handlers.HelloHandler)

	return router
}
