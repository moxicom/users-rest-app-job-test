package services

import (
	"log/slog"

	"github.com/moxicom/user_test/internal/models"
	"github.com/moxicom/user_test/internal/storage"
)

type User interface {
	GetUsers(models.Filters) ([]models.User, error)
	CreateUser(string) (uint, error)
	DeleteUser(int) error
	UpdateUser(models.User) error
}

type Tasks interface {
	GetUserTasks(userID int)
}

type Service struct {
	User
}

func New(s storage.Storage, log *slog.Logger) *Service {
	return &Service{
		User: newUserService(s, log),
	}
}
