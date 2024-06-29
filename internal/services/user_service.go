package services

import (
	"log/slog"

	"github.com/moxicom/user_test/internal/models"
	"github.com/moxicom/user_test/internal/storage"
)

type UserService struct {
	s   storage.Storage
	log *slog.Logger
}

func newUserService(s storage.Storage, log *slog.Logger) *UserService {
	return &UserService{s, log}
}

func (s *UserService) CreateUser(passport string) (uint, error) {
	log := s.log.With(slog.String("op", "service.CreateUser"))

	// TODO: UNCOMMIT LATER - use api to get user data
	// user, err := utils.GetUserData(passport)
	// if err != nil {
	// 	log.Error("failed to get user data")
	// 	return 0, err
	// }

	user := models.User{
		PassportNumber: passport,
		Surname:        "alexio" + passport,
		Name:           "alex" + passport,
		Patronymic:     "alexovich" + passport,
		Address:        "alex street" + passport,
	}

	log.Debug("Got new user info", slog.Any("user", user))

	userID, err := s.s.AddUser(user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (s *UserService) GetUsers(f models.Filters) ([]models.User, error) {
	return s.s.GetUsers(f)
}

func (s *UserService) DeleteUser(userID uint) error {
	return s.s.DeleteUser(userID)
}

func (s *UserService) UpdateUser(userID uint, filters models.Filters) error {
	return s.s.UpdateUser(userID, filters)
}
