package models

type UserFilters struct {
	PassportNumber string
	Surname        string
	Name           string
	Patronymic     string
	Address        string
}

type TaskFilters struct {
	Asc bool
}
