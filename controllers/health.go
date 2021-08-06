package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health check Pattern. Service status.
func Health(c *gin.Context) {

	// the status of the connections to the infrastructure services used by the service instance
	// the status of the host, e.g. disk space
	// application specific logic
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
