package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moxicom/user_test/internal/models"
	"github.com/moxicom/user_test/internal/utils"
)

type createUser struct {
	PassportNumber string `json:"passportNumber" binding:"required"`
}

func (h *Handler) CreateUser(c *gin.Context) {
	log := h.log.With(slog.String("op", "handler.CreateUser"))

	var user createUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	ok := utils.ValidatePassword(user.PassportNumber)
	if !ok {
		log.Error(fmt.Sprintf("invalid passport number: %s", user.PassportNumber))
		c.JSON(http.StatusBadRequest, Message{"invalid passport number"})
		return
	}

	userID, err := h.service.User.CreateUser(user.PassportNumber)
	if err != nil {
		log.Error("failed to create user")
		c.JSON(http.StatusInternalServerError, Message{"failed to create user"})
		return
	}

	c.JSON(http.StatusOK, Message{fmt.Sprintf("%d", userID)})
}

func (h *Handler) GetUsers(c *gin.Context) {
	filters := models.Filters{
		PassportNumber: c.Query("passport_number"),
		Surname:        c.Query("surname"),
		Name:           c.Query("name"),
		Patronymic:     c.Query("patronymic"),
		Address:        c.Query("address"),
	}

	// TODO: service - get users
	users, err := h.service.GetUsers(filters)
	if err != nil {
		h.log.Error("failed to get users")
		c.JSON(http.StatusInternalServerError, Message{"failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	f := models.Filters{
		PassportNumber: c.Query("passport_number"),
		Surname:        c.Query("surname"),
		Name:           c.Query("name"),
		Patronymic:     c.Query("patronymic"),
		Address:        c.Query("address"),
	}
	c.Param("id")

	if f.Address == "" && f.Name == "" && f.Patronymic == "" && f.PassportNumber == "" && f.Surname == "" {
		c.JSON(http.StatusBadRequest, Message{"no data to update. use passport_number, surname, name, patronymic, address"})
		return
	}

	// Validate password number
	if f.PassportNumber != "" {
		ok := utils.ValidatePassword(f.PassportNumber)
		if !ok {
			h.log.Error(fmt.Sprintf("invalid passport number: %s", f.PassportNumber))
			c.JSON(http.StatusBadRequest, Message{"invalid passport number"})
			return
		}
	}

	// TODO: service - update user
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	h.log.Debug(id)
	// TODO: service - delete User
}

func (h *Handler) GetUsersWithTasks(c *gin.Context) {
	// TODO: get users tasks
}
