package services

import (
	"log/slog"
	"time"

	"github.com/moxicom/user_test/internal/models"
	"github.com/moxicom/user_test/internal/storage"
)

type TaskService struct {
	s   storage.Storage
	log *slog.Logger
}

func newTaskService(s storage.Storage, log *slog.Logger) *TaskService {
	return &TaskService{s, log}
}

func (s *TaskService) CreateTask(task models.Task) (uint, error) {
	task.StartTime = time.Now()
	return s.s.CreateTask(task)
}

func (s *TaskService) EndTask(taskID uint) error {
	endTime := time.Now()
	return s.s.EndTask(taskID, endTime)
}

func (s *TaskService) DeleteTask(taskID uint) error {
	return s.s.DeleteTask(taskID)
}
