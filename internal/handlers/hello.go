package handlers

import (
	"net/http"

	"golang-hello-world/internal/jokes"

	"github.com/gin-gonic/gin"
)

func HelloHandler(context *gin.Context) {
	var joke string = jokes.GetRandomJoke()

	context.JSON(http.StatusOK, gin.H{
		"message": "Hello, world!",
		"joke":    joke,
	})
}
