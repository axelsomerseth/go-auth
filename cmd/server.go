package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.RedirectTrailingSlash = true

	// Attach middlewares to router.
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	router.Run()
}
