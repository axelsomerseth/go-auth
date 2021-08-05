package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		router *gin.Engine
	)

	router = gin.New()
	router.RedirectTrailingSlash = true

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	router.Run()
}
