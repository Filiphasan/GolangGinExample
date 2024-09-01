package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	PostgresConnection string
	RedisConnection    string
	ElasticUser        string
	ElasticPassword    string
}

var appConfig AppConfig

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")

	postgresConnection := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=go-gin-example", postgresHost, postgresUser, postgresPassword)

	appConfig = AppConfig{
		PostgresConnection: postgresConnection,
	}

	log.Println("App Config: ", appConfig)
}

func GetAppConfig() AppConfig {
	return appConfig
}
