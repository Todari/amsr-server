package main

import (
	"net/http"

	"github.com/Todari/amsr-server/configs"
	"github.com/Todari/amsr-server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Connect Database
	configs.ConnectDB()

	// Hello, World!
	router.GET("/", func(ctx_ *gin.Context) {
		ctx_.String(http.StatusOK, "Hello, World!")
	})

	protected := router.Group("/")
	routes.UserRouter(protected)

	err := router.Run("0.0.0.0:8080")

	if err != nil {
		return
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Origin", "http://amsr.site")
		// ctx.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		ctx.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}
