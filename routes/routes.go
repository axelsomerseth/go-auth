package routes

import (
	"github.com/axelsomerseth/go-auth/handlers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {

	// Health check pattern.
	router.GET("/health", handlers.Health)

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			// Auth routes.
			v1.POST("/signup", handlers.Register)
			v1.POST("/login", handlers.Login)
			v1.POST("/logout", handlers.Logout)
			v1.GET("/user", handlers.User)
		}
	}
}
