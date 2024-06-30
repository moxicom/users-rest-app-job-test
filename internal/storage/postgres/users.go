package postgres

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/moxicom/user_test/internal/models"
)

func (p *PgStorage) AddUser(user models.User) (uint, error) {
	log := p.log.With(slog.String("op", "PgStorage.AddUser"))

	result := p.db.Create(&user)
	if result.Error != nil {
		log.Error("failed to add user", result.Error)
		return 0, result.Error
	}
	log.Debug("user added to storage", slog.Any("user", user))
	return user.ID, nil
}

func (p *PgStorage) GetUsers(filters models.UserFilters) ([]models.User, error) {
	log := p.log.With(slog.String("op", "PgStorage.GetUsers"))
	var users []models.User

	tx := p.db.Begin()

	query := tx.Model(&models.User{})

	if filters.PassportNumber != "" {
		query = query.Where("LOWER(passport_number) LIKE LOWER(?)", "%"+filters.PassportNumber+"%")
	}
	if filters.Surname != "" {
		query = query.Where("LOWER(surname) LIKE LOWER(?)", "%"+filters.Surname+"%")
	}
	if filters.Name != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+filters.Name+"%")
	}
	if filters.Patronymic != "" {
		query = query.Where("LOWER(patronymic) LIKE LOWER(?)", "%"+filters.Patronymic+"%")
	}
	if filters.Address != "" {
		query = query.Where("LOWER(address) LIKE LOWER(?)", "%"+filters.Address+"%")
	}

	res := query.Find(&users)
	if res.Error != nil {
		log.Error("failed to get users", res.Error)
		return []models.User{}, res.Error
	}

	log.Debug("users found", slog.Any("users", res.RowsAffected))

	return users, tx.Commit().Error
}

func (p *PgStorage) UpdateUser(userID uint, filters models.UserFilters) error {
	log := p.log.With(slog.String("op", "PgStorage.UpdateUser"))
	var user models.User

	tx := p.db.Begin()
	defer tx.Rollback()

	if err := tx.First(&user, userID).Error; err != nil {
		return err
	}

	// Update fields based on non-empty filter values
	if filters.PassportNumber != "" {
		user.PassportNumber = filters.PassportNumber
	}
	if filters.Surname != "" {
		user.Surname = filters.Surname
	}
	if filters.Name != "" {
		user.Name = filters.Name
	}
	if filters.Patronymic != "" {
		user.Patronymic = filters.Patronymic
	}
	if filters.Address != "" {
		user.Address = filters.Address
	}

	if err := tx.Save(&user).Error; err != nil {
		log.Error("failed to update user", err)
		return err
	}

	return tx.Commit().Error
}

func (p *PgStorage) DeleteUser(userID uint) error {
	log := p.log.With(slog.String("op", "PgStorage.DeleteUser"))
	tx := p.db.Begin()

	res := tx.Delete(&models.User{}, userID)
	if res.Error != nil {
		log.Error("failed to delete user. Rolled back", res.Error)
		return res.Error
	}

	return tx.Commit().Error
}

func (p *PgStorage) GetUserTasks(userID uint, startTime time.Time, endTime time.Time, isAsc bool) ([]models.TaskWithTotalTime, error) {
	log := p.log.With(slog.String("op", "PgStorage.GetUserTasks"))

	var tasks []models.TaskWithTotalTime
	tx := p.db.Begin()
	defer tx.Rollback()

	db := p.db

	subquery := db.Model(&models.TaskPeriod{}).
		Select("task_id, SUM(EXTRACT(EPOCH FROM COALESCE(end_time, CURRENT_TIMESTAMP) - COALESCE(start_time, CURRENT_TIMESTAMP))) AS total_duration").
		Group("task_id")

	// Main query to fetch tasks with total durations
	res := db.
		// Joins("LEFT JOIN (?) AS periods ON tasks.id = periods.task_id", subquery).
		// Model(&models.Task{}).
		// Select("tasks.*, COALESCE(periods.total_duration, 0) AS duration").
		// Where("tasks.user_id = ? AND tasks.created_at BETWEEN ? AND ?", userID, startTime, endTime).
		// Order("duration DESC").
		// Find(&tasks)
		Joins("LEFT JOIN (?) AS periods ON tasks.id = periods.task_id", subquery).
		Model(&models.Task{}).
		Select(`
        tasks.*,
        COALESCE(periods.total_duration, 0) AS total_seconds,
        FLOOR(COALESCE(periods.total_duration, 0) / 3600) AS duration_hours,
        FLOOR(COALESCE(periods.total_duration, 0) / 60) AS duration_minutes
    `).
		Where("tasks.user_id = ? AND tasks.created_at BETWEEN ? AND ?", userID, startTime, endTime).
		Order("total_seconds DESC").
		Find(&tasks)

	// res := subquery.Find(&tasks)

	err := res.Error
	if err != nil {
		log.Error("Failed to get user tasks", slog.Uint64("user_id", uint64(userID)), slog.Any("err", err.Error()))
		return nil, err
	}

	// Convert total_time from seconds to time.Duration
	// for i := range tasks {
	// 	tasks[i].Duration = time.Duration(tasks[i].Duration) * time.Second
	// }
	fmt.Println()
	fmt.Println(tasks)
	fmt.Println()
	return tasks, nil
}
