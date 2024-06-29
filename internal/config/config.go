package config

import (
	"os"

	"github.com/moxicom/user_test/internal/storage/postgres"
)

func InitDbConfig() postgres.PgConfig {
	return postgres.PgConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Dbname:   os.Getenv("POSTGRES_DB"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("SSL_MODE"),
	}
}
