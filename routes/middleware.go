package routes

import (
	"fmt"
	"kryptonim/app/container"
	"kryptonim/app/handlers"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cont *container.Container) *gin.Engine {
	r := gin.Default()

	// Middleware for token-based authorization
	r.Use(AuthMiddleware())

	fmt.Println()
	r.GET("/rates", handlers.GetRates(cont))
	r.GET("/exchange", handlers.GetExchange(cont))

	return r
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authReq := os.Getenv("AUTH_REQ")
		if authReq == "OFF" {
			c.Next()
			return
		}

		token := c.GetHeader("Auth")
		if token != os.Getenv("AUTH_TOKEN") {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
