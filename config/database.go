package config

import (
	"go-gin-example/src/model/entity"
	"go-gin-example/src/seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB(appConfig *AppConfig) (*gorm.DB, error) {
	database, gormErr := gorm.Open(postgres.Open(appConfig.PostgresConnection), &gorm.Config{})
	if gormErr != nil {
		log.Fatal("Error connecting to the database: ", gormErr)
		return nil, gormErr
	}

	migrationErr := database.AutoMigrate(
		&entity.User{},
		&entity.Permission{},
		&entity.UserPermission{},
	)
	if migrationErr != nil {
		log.Fatal("Error migrating the database: ", migrationErr)
		return nil, migrationErr
	}

	seed.CreateBasePermissions(database)

	return database, nil
}
