package database

import (
	"fmt"

	"github.com/axelsomerseth/go-auth/models"

	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Connect(databaseDSN string) error {
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
	db, err = gorm.Open(postgres.Open(databaseDSN), gormConfig)
	if err != nil {
		return err
	}

	// Automigrate option
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	fmt.Printf("database connected: %s \n", db.Migrator().CurrentDatabase())

	return nil
}
