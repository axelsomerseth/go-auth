package main

import (
	"log"

	"github.com/axelsomerseth/go-auth/database"
	"github.com/axelsomerseth/go-auth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect database storage.
	if err := database.Connect("host=database port=5432 user=postgres password=password dbname=GO_AUTH_DB sslmode=disable"); err != nil {
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
