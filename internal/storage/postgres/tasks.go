package postgres

import (
	"log/slog"
	"time"

	"github.com/moxicom/user_test/internal/models"
	"github.com/moxicom/user_test/internal/storage"
)

func (p *PgStorage) CreateTask(task models.Task) (uint, error) {
	log := p.log.With(slog.String("op", "PgStorage.CreateTask"))

	result := p.db.Create(&task)
	if result.Error != nil {
		log.Error("failed to add task", slog.Any("err", result.Error))
		return 0, result.Error
	}
	log.Debug("user added to storage", slog.Any("task", task))
	return task.ID, nil
}

func (p *PgStorage) FinishTask(taskID uint, finishTime time.Time) error {
	log := p.log.With(slog.String("op", "PgStorage.EndTask"))
	// TODO - fix
	// TODO - end period

	err := p.EndPeriod(taskID, finishTime)

	if err != storage.ErrPeriodNotStarted {
		return err
	}

	var task models.Task

	tx := p.db.Begin()
	defer tx.Rollback()

	if err := tx.First(&task, taskID).Error; err != nil {
		log.Error("Error selecting task on ending ", slog.Any("err", err))
		return err
	}

	task.IsFinished = true

	if err := tx.Save(&task).Error; err != nil {
		log.Error("failed to finish task", slog.Any("err", err))
		return err
	}

	return tx.Commit().Error
}

func (p *PgStorage) DeleteTask(taskID uint) error {
	log := p.log.With(slog.String("op", "PgStorage.DeleteTask"))
	tx := p.db.Begin()
	defer tx.Rollback()

	res := tx.Delete(&models.Task{}, taskID)
	if res.Error != nil {
		log.Error("failed to delete task. Rolled back", slog.Any("err", res.Error))
		return res.Error
	}

	return tx.Commit().Error
}

func (p *PgStorage) StartPeriod(taskID uint, startTime time.Time) error {
	log := p.log.With(slog.String("op", "PgStorage.StartPeriod"))

	var ongoingPeriod models.TaskPeriod
	tx := p.db.Begin()
	defer tx.Rollback()

	tx.Where("task_id = ? AND end_time IS NULL", taskID).Last(&ongoingPeriod)
	if ongoingPeriod.ID != 0 {
		log.Warn("task can not be started. Should be finished", slog.Any("err", storage.ErrPeriodNotFinished))
		return storage.ErrPeriodNotFinished
	}

	now := time.Now()
	period := models.TaskPeriod{TaskID: taskID, StartTime: &now}
	res := tx.Create(&period)
	if res.Error != nil {
		log.Error("failed to start period", slog.Uint64("task_id", uint64(taskID)), slog.Any("err", res.Error))
		return res.Error
	}

	return tx.Commit().Error
}

func (p *PgStorage) EndPeriod(taskID uint, startTime time.Time) error {
	log := p.log.With(slog.String("op", "PgStorage.EndPeriod"))

	var ongoingPeriod models.TaskPeriod
	tx := p.db.Begin()
	defer tx.Rollback()

	tx.Where("task_id = ? AND end_time IS NULL", taskID).Last(&ongoingPeriod)
	if ongoingPeriod.ID == 0 {
		log.Warn("task can not be finished. Should be started", slog.Any("err", storage.ErrPeriodNotFinished))
		return storage.ErrPeriodNotStarted
	}

	now := time.Now()
	ongoingPeriod.EndTime = &now

	res := tx.Save(&ongoingPeriod)
	if res.Error != nil {
		log.Error("failed to end period", slog.Uint64("task_id", uint64(taskID)), slog.Any("err", res.Error))
		return res.Error
	}
	return tx.Commit().Error
}
