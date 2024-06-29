package models

type User struct {
	ID             uint   `gorm:"primarykey"`
	PassportNumber string `gorm:"uniqueIndex" json:"passport_number"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}

type Task struct {
	ID       uint `gorm:"primarykey"`
	UserID   uint
	TaskName string
	Duration float64
}
