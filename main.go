package main

import (
	"log"

	"github.com/axelsomerseth/go-auth/database"
	"github.com/axelsomerseth/go-auth/internal/config"
	"github.com/axelsomerseth/go-auth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	appConfig, err := config.Init()
	if err != nil {
		log.Fatalf("error trying to set up server configuration and env vars: ", err)
	}

	// Connect database storage.
	if err := database.Connect(appConfig); err != nil {
		log.Fatalf("error trying to connect the database, due %s", err)
	}

	router := gin.New()
	router.RedirectTrailingSlash = true

	// Attach middlewares to router.
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.SetRoutes(router)

	router.Run()
}
