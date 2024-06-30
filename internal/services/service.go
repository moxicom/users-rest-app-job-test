package services

import (
	"log/slog"
	"time"

	"github.com/moxicom/user_test/internal/models"
	"github.com/moxicom/user_test/internal/storage"
)

type User interface {
	GetUsers(models.UserFilters) ([]models.User, error)
	GetUserTasks(uint, time.Time, time.Time, models.TaskFilters) ([]models.TaskWithTotalTime, error)
	CreateUser(string) (uint, error)
	DeleteUser(uint) error
	UpdateUser(uint, models.UserFilters) error
}

type Task interface {
	CreateTask(models.Task) (uint, error)
	FinishTask(uint) error
	DeleteTask(uint) error
	StartPeriod(uint) error
	EndPeriod(uint) error
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
