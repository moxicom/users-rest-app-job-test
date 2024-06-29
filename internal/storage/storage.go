package storage

import "github.com/moxicom/user_test/internal/models"

type Storage interface {
	GetUsers(f models.Filters) ([]models.User, error)
	AddUser(models.User) (uint, error)
	UpdateUser(uint, models.Filters) error
	DeleteUser(userID uint) error
}
