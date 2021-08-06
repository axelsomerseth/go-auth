package routes

import (
	"github.com/axelsomerseth/go-auth/controllers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {

	// Health check pattern.
	router.GET("/health", controllers.Health)

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			// Auth routes.
			v1.POST("/signup", controllers.Register)
			v1.POST("/login", controllers.Login)
			v1.POST("/logout", controllers.Logout)
			v1.GET("/user", controllers.User)
		}
	}
}
