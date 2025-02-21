package handlers

import (
	"kryptonim/app/container"

	"github.com/gin-gonic/gin"
)

func GetExchange(container *container.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		container.ExchangeService.HandleExchange(c)
	}
}
