package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const databaseEnvKey = "DATABASE_URL"

func SetupModels(dbConn string) *gorm.DB {
	// Enable viper to read environment variables

	// dbConn := fmt.Sprintf("%v", viper.Get(databaseEnvKey))

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{})

	return db
}
