package postgres

import (
	"fmt"
	"log"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgConfig struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
	SSLMode  string
}

type PgStorage struct {
	db  *gorm.DB
	log *slog.Logger
}

func NewDbInit(cfg PgConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Dbname,
		cfg.Port,
		cfg.SSLMode)

	log.Println("dsn", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewStorage(db *gorm.DB, log *slog.Logger) *PgStorage {
	return &PgStorage{db, log}
}
