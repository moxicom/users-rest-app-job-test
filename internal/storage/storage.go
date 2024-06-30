package storage

import (
	"fmt"
	"time"

	"github.com/moxicom/user_test/internal/models"
)

var (
	ErrPeriodNotStarted  = fmt.Errorf("period not started")
	ErrPeriodNotFinished = fmt.Errorf("period not finished")
)

type Storage interface {
	GetUsers(models.UserFilters) ([]models.User, error)
	GetUserTasks(userID uint, startTime time.Time, endTime time.Time, isAsc bool) ([]models.TaskWithTotalTime, error)
	AddUser(models.User) (uint, error)
	UpdateUser(uint, models.UserFilters) error
	DeleteUser(uint) error

	CreateTask(models.Task) (uint, error)
	FinishTask(uint, time.Time) error
	DeleteTask(uint) error
	StartPeriod(uint, time.Time) error
	EndPeriod(uint, time.Time) error
}
