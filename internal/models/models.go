package models

import "time"

type User struct {
	ID             uint   `gorm:"primarykey"`
	PassportNumber string `gorm:"uniqueIndex" json:"passport_number"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
	Tasks          []Task `json:"-" gorm:"constraint:OnDelete:CASCADE;"` // Establish the relationship and enable cascading deletes
}

type Task struct {
	ID         uint         `json:"id" gorm:"primarykey"`
	UserID     uint         `json:"user_id" binding:"required" gorm:"index"`
	TaskName   string       `json:"task_name" binding:"required"`
	CreatedAt  time.Time    `json:"created_at"`
	IsFinished bool         `json:"is_finished"`
	Periods    []TaskPeriod `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
}

type TaskWithTotalTime struct {
	Task
	TotalSeconds    string `json:"-"`
	DurationHours   int    `json:"duration_hours"`
	DurationMinutes int    `json:"duration_minutes"`
}

type TaskPeriod struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	TaskID    uint       `json:"task_id" gorm:"index"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
}
