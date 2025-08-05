package main

import (
	"golang-hello-world/internal/routes"
)

func main() {
	var app = routes.SetupRouter()
	app.Run(":8080")
}
