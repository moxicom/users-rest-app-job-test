package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/moxicom/user_test/internal/config"
	"github.com/moxicom/user_test/internal/handlers"
	"github.com/moxicom/user_test/internal/server"
	"github.com/moxicom/user_test/internal/services"
	"github.com/moxicom/user_test/internal/storage/migrations"
	"github.com/moxicom/user_test/internal/storage/postgres"
	"github.com/moxicom/user_test/internal/utils"
)

var (
	envLog string
)

//	@title			time-tracker application
//	@version		0.1
//	@description	This is a simple backend for time-tracker application without authorization

// @BasePath	/
func main() {
	runServer(context.Background())
	log.Println("DB_HOST:", os.Getenv("DB_HOST"))
	log.Println("POSTGRES_USER:", os.Getenv("POSTGRES_USER"))
	log.Println("POSTGRES_PASSWORD:", os.Getenv("POSTGRES_PASSWORD"))
	log.Println("POSTGRES_DB:", os.Getenv("POSTGRES_DB"))
	log.Println("DB_PORT:", os.Getenv("DB_PORT"))
	log.Println("SSL_MODE:", os.Getenv("SSL_MODE"))
}

func runServer(ctx context.Context) error {
	flag.StringVar(
		&envLog,
		"envLog",
		utils.EnvProd,
		fmt.Sprintf("'%s' or '%s' to setup logger", utils.EnvProd, utils.EnvLocal),
	)
	log := utils.SetupLogger(envLog)

	if err := godotenv.Load(); err != nil {
		log.Error("%s", slog.Any("err", err))
		return err
	}

	apiAddress := os.Getenv("API_ADDRESS")
	if apiAddress == "" {
		return fmt.Errorf("API_ADDRESS is not set")
	}
	utils.ApiAddress = apiAddress

	cfg := config.InitDbConfig()

	db, err := postgres.NewDbInit(cfg)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	migrations.MigratePostgres(db, log)

	// Dependency injection
	storage := postgres.NewStorage(db, log)
	service := services.New(storage, log)
	handler := handlers.New(service, log)
	server := server.New()

	go func() {
		if err = server.Run(os.Getenv("SERVER_PORT"), handler.InitRoutes()); err != nil {
			log.Error("listen and serve: %s", slog.Any("err", err))
			return
		}
	}()

	<-ctx.Done()

	log.Info("Shutting down gracefully")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Close(); err != nil {
		return err
	}

	return nil
}
