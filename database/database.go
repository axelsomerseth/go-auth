package database

import (
	"fmt"
	"log"

	"github.com/axelsomerseth/go-auth/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Connect() {
	var (
		err        error
		gormConfig *gorm.Config
	)

	// Define gorm configuration.
	gormConfig = &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	}

	// Open connection.
	DB, err = gorm.Open(sqlite.Open("/Users/axelsomerseth/GO-AUTH.db"), gormConfig)
	if err != nil {
		log.Fatalf("error connecting to the database, due %s", err)
	}

	// Automigrate option
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("cannot automigrate database, due %s", err)
	}

	fmt.Print("database connected")
}
