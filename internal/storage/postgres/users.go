package postgres

import (
	"log/slog"

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

func (p *PgStorage) GetUsers(filters models.Filters) ([]models.User, error) {
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

func (p *PgStorage) UpdateUser(userID uint, filters models.Filters) error {
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
