package models

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
	ID       uint `gorm:"primarykey"`
	UserID   uint `gorm:"index"`
	TaskName string
	Duration float64
}
