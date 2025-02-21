package errors

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, statusCode int) {
	log.Printf("Error: %v", err)
	c.JSON(statusCode, gin.H{"error": err.Error()})
	c.Abort()
}
