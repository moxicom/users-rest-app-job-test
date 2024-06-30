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
	task.CreatedAt = time.Now()
	task.IsFinished = false
	return s.s.CreateTask(task)
}

func (s *TaskService) FinishTask(taskID uint) error {
	endTime := time.Now()
	return s.s.FinishTask(taskID, endTime)
}

func (s *TaskService) DeleteTask(taskID uint) error {
	return s.s.DeleteTask(taskID)
}

func (s *TaskService) StartPeriod(taskID uint) error {
	return s.s.StartPeriod(taskID, time.Now())
}

func (s *TaskService) EndPeriod(taskID uint) error {
	return s.s.EndPeriod(taskID, time.Now())
}
