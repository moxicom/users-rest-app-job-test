package services

import (
	"log/slog"

	"github.com/moxicom/user_test/internal/models"
	"github.com/moxicom/user_test/internal/storage"
)

type User interface {
	GetUsers(models.Filters) ([]models.User, error)
	CreateUser(string) (uint, error)
	DeleteUser(uint) error
	UpdateUser(uint, models.Filters) error
}

type Task interface {
	CreateTask(models.Task) (uint, error)
	EndTask(uint) error
	DeleteTask(uint) error
}

type Service struct {
	Task
	User
}

func New(s storage.Storage, log *slog.Logger) *Service {
	return &Service{
		User: newUserService(s, log),
		Task: newTaskService(s, log),
	}
}
