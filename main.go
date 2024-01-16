package main

import (
	"net/http"

	"github.com/Todari/amsr-server/configs"
	"github.com/Todari/amsr-server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect Database
	configs.ConnectDB()

	// Hello, World!
	router.GET("/", func(ctx_ *gin.Context) {
		ctx_.String(http.StatusOK, "Hello, World!")
	})

	// public := router.Group("/")

	protected := router.Group("/")
	routes.UserRouter(protected)

	err := router.Run("0.0.0.0:8080")

	if err != nil {
		return
	}
}
