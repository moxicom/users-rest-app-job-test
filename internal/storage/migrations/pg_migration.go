package migrations

import (
	"log/slog"

	"github.com/moxicom/user_test/internal/models"
	"gorm.io/gorm"
)

func MigratePostgres(db *gorm.DB, log *slog.Logger) {
	log.Info("Making automigration...")
	db.AutoMigrate(&models.User{}, &models.Task{})
}
