package postgres

import (
	"log/slog"
	"time"

	"github.com/moxicom/user_test/internal/models"
)

func (p *PgStorage) CreateTask(task models.Task) (uint, error) {
	log := p.log.With(slog.String("op", "PgStorage.CreateTask"))

	result := p.db.Create(&task)
	if result.Error != nil {
		log.Error("failed to add task", result.Error)
		return 0, result.Error
	}
	log.Debug("user added to storage", slog.Any("task", task))
	return task.ID, nil
}

func (p *PgStorage) EndTask(taskID uint, endTime time.Time) error {
	log := p.log.With(slog.String("op", "PgStorage.EndTask"))
	var task models.Task

	tx := p.db.Begin()
	defer tx.Rollback()

	if err := tx.First(&task, taskID).Error; err != nil {
		log.Error("Error selecting task on ending ", err)
		return err
	}

	task.EndTime = endTime

	if err := tx.Save(&task).Error; err != nil {
		log.Error("failed to update task", err)
		return err
	}

	return tx.Commit().Error
}

func (p *PgStorage) DeleteTask(taskID uint) error {
	log := p.log.With(slog.String("op", "PgStorage.DeleteTask"))
	tx := p.db.Begin()

	res := tx.Delete(&models.Task{}, taskID)
	if res.Error != nil {
		log.Error("failed to delete task. Rolled back", res.Error)
		return res.Error
	}

	return tx.Commit().Error
}
