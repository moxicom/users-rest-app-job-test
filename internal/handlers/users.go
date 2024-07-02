package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moxicom/user_test/internal/models"
	"github.com/moxicom/user_test/internal/utils"
)

const (
	asc  = "asc"
	desc = "desc"
)

type createUser struct {
	PassportNumber string `json:"passportNumber" binding:"required"`
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided passport number
// @Tags users
// @Accept json
// @Produce json
// @Param user body createUser true "User"
// @Success 200 {object} Message "User ID"
// @Failure 400 {object} Message "Invalid body data or invalid passport number"
// @Failure 500 {object} Message "Failed to create user"
// @Router /users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	log := h.log.With(slog.String("op", "handler.CreateUser"))

	var user createUser
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error("error while parsing json ", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, Message{"invalid body data"})
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

	log.Info("User created successfully", slog.Uint64("user_id", uint64(userID)))
	c.JSON(http.StatusOK, Message{fmt.Sprintf("%d", userID)})
}

// GetUsers retrieves users based on filters
// @Summary Get users
// @Description Retrieve users based on filters
// @Tags users
// @Accept json
// @Produce json
// @Param passport_number query string false "Passport Number"
// @Param surname query string false "Surname"
// @Param name query string false "Name"
// @Param patronymic query string false "Patronymic"
// @Param address query string false "Address"
// @Success 200 {array} models.User "List of users"
// @Failure 500 {object} Message "Failed to get users"
// @Router /users [get]
func (h *Handler) GetUsers(c *gin.Context) {
	log := h.log.With(slog.String("op", "handler.GetUsers"))
	filt := utils.GetFilters(c)

	users, err := h.service.User.GetUsers(filt)
	if err != nil {
		log.Error("failed to get users")
		c.JSON(http.StatusInternalServerError, Message{"failed to get users"})
		return
	}

	log.Info("User found successfully")
	c.JSON(http.StatusOK, users)
}

// UpdateUser updates a user
// @Summary Update a user
// @Description Update a user with the provided data
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param passport_number query string false "Passport Number"
// @Param surname query string false "Surname"
// @Param name query string false "Name"
// @Param patronymic query string false "Patronymic"
// @Param address query string false "Address"
// @Success 200 {object} Message "User updated"
// @Failure 400 {object} Message "Incorrect ID or invalid input data"
// @Failure 500 {object} Message "Failed to update user"
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	log := h.log.With(slog.String("op", "handler.UpdateUser"))
	filt := utils.GetFilters(c)
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Warn("Failed to parse user ID", slog.String("id", c.Param("id")), slog.Any("err", err))
		c.JSON(http.StatusBadRequest, Message{"incorrect id"})
		return
	}

	if filt.Address == "" && filt.Name == "" && filt.Patronymic == "" && filt.PassportNumber == "" && filt.Surname == "" {
		log.Warn("No data to update for user", slog.Uint64("user_id", id64))
		c.JSON(http.StatusBadRequest, Message{"no data to update. use passport_number, surname, name, patronymic, address"})
		return
	}

	// Validate password number
	if filt.PassportNumber != "" {
		ok := utils.ValidatePassword(filt.PassportNumber)
		if !ok {
			log.Warn("Invalid passport number", slog.String("passport_number", filt.PassportNumber))
			c.JSON(http.StatusBadRequest, Message{"invalid passport number"})
			return
		}
	}

	err = h.service.User.UpdateUser(uint(id64), filt)
	if err != nil {
		log.Error("Failed to update user")
		c.JSON(http.StatusInternalServerError, Message{"failed to update user"})
		return

	}

	log.Info("User updated successfully", slog.Uint64("user_id", id64))
	c.JSON(http.StatusOK, Message{"user updated"})
}

// DeleteUser deletes a user
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} Message "User deleted"
// @Failure 400 {object} Message "ID should be an integer"
// @Failure 500 {object} Message "Failed to delete user"
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	log := h.log.With(slog.String("op", "handler.DeleteUser"))
	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Warn("Invalid user ID format", slog.String("id", id), slog.Any("err", err))
		c.JSON(http.StatusBadRequest, Message{"id should be integer"})
		return
	}

	err = h.service.User.DeleteUser(uint(id64))
	if err != nil {
		log.Error("Failed to delete user", slog.String("id", id))
		c.JSON(http.StatusInternalServerError, Message{"failed to delete user"})
		return
	}

	log.Info("User deleted successfully", slog.String("id", id))
	c.JSON(http.StatusOK, Message{"user deleted"})
}

// GetUsersWithTasks gets the tasks for a user
// @Summary Get user tasks
// @Description Get tasks for a user within a specified date range and with optional sorting
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param start_date query string true "Start date in RFC3339 format"
// @Param end_date query string true "End date in RFC3339 format"
// @Param sort query string false "Sort order, can be 'asc' or 'desc'"
// @Success 200 {array} models.Task "Tasks found"
// @Failure 400 {object} Message "Invalid input data"
// @Failure 500 {object} Message "Failed to get tasks for user"
// @Router /users/{id}/tasks [get]
func (h *Handler) GetUsersWithTasks(c *gin.Context) {
	log := h.log.With(slog.String("op", "handler.GetUsersWithTasks"))
	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Warn("Invalid user ID format", slog.String("id", id), slog.Any("err", err))
		c.JSON(http.StatusBadRequest, Message{"id should be integer"})
		return
	}

	filters := models.TaskFilters{}
	if sortFilter := c.Query("sort"); sortFilter == asc {
		filters.Asc = true
	} else {
		filters.Asc = false
	}

	startDate, err := time.Parse(time.RFC3339, c.Query("start_date"))
	if err != nil {
		log.Warn("No start date to get tasks for user", slog.Uint64("user_id", id64), slog.Any("err", err))
		c.JSON(http.StatusBadRequest, Message{"Invalid start date"})
		return
	}

	endDate, err := time.Parse(time.RFC3339, c.Query("end_date"))
	if err != nil {
		log.Warn("No end date to get tasks for user", slog.Uint64("user_id", id64))
		c.JSON(http.StatusBadRequest, Message{"Invalid end date"})
		return
	}

	tasks, err := h.service.User.GetUserTasks(uint(id64), startDate, endDate, filters)
	if err != nil {
		log.Error("Failed to get tasks for user", slog.Uint64("user_id", id64))
		c.JSON(http.StatusInternalServerError, Message{"Failed to get tasks for user"})
		return
	}

	log.Info("Successfully found tasks for user", slog.Uint64("user_id", id64))
	c.JSON(http.StatusOK, tasks)
}
