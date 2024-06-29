package storage

import (
	"time"

	"github.com/moxicom/user_test/internal/models"
)

type Storage interface {
	GetUsers(models.Filters) ([]models.User, error)
	AddUser(models.User) (uint, error)
	UpdateUser(uint, models.Filters) error
	DeleteUser(uint) error

	CreateTask(models.Task) (uint, error)
	EndTask(uint, time.Time) error
	DeleteTask(uint) error
}
