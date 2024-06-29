package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	log := h.log.With(slog.String("op", "handler.GetUsers"))
	filt := utils.GetFilters(c)

	users, err := h.service.GetUsers(filt)
	if err != nil {
		log.Error("failed to get users")
		c.JSON(http.StatusInternalServerError, Message{"failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	log := h.log.With(slog.String("op", "handler.UpdateUser"))
	filt := utils.GetFilters(c)
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Error("failed to parse id")
		c.JSON(http.StatusBadRequest, Message{"incorrect id"})
		return
	}

	if filt.Address == "" && filt.Name == "" && filt.Patronymic == "" && filt.PassportNumber == "" && filt.Surname == "" {
		log.Error("no data to update")
		c.JSON(http.StatusBadRequest, Message{"no data to update. use passport_number, surname, name, patronymic, address"})
		return
	}

	// Validate password number
	if filt.PassportNumber != "" {
		ok := utils.ValidatePassword(filt.PassportNumber)
		if !ok {
			h.log.Error(fmt.Sprintf("invalid passport number: %s", filt.PassportNumber))
			c.JSON(http.StatusBadRequest, Message{"invalid passport number"})
			return
		}
	}

	err = h.service.UpdateUser(uint(id64), filt)
	if err != nil {
		log.Error("failed to update user", err)
		c.JSON(http.StatusInternalServerError, Message{"failed to update user"})
		return

	}

	c.JSON(http.StatusOK, Message{"user updated"})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	log := h.log.With(slog.String("op", "handler.DeleteUser"))
	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	err = h.service.DeleteUser(uint(id64))
	if err != nil {
		log.Error("failed to delete user", slog.String("id", id))
		c.JSON(http.StatusInternalServerError, Message{"failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, Message{"user deleted"})
}

func (h *Handler) GetUsersWithTasks(c *gin.Context) {
	// TODO: get users tasks
}
