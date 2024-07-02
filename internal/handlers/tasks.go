package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moxicom/user_test/internal/models"
	"github.com/moxicom/user_test/internal/storage"
)

// CreateTask creates a new task
// @Summary Create a new task
// @Description Create a new task for a user
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Task object"
// @Success 200 {object} Message "Task created successfully"
// @Failure 400 {object} Message "Invalid body data"
// @Failure 500 {object} Message "Failed to create task"
// @Router /tasks [post]
func (h *Handler) CreateTask(c *gin.Context) {
	log := h.log.With(slog.String("op", "Handler.CreateTask"))
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Error("error while parsing json ", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, Message{"invalid body data"})
		return
	}

	taskID, err := h.service.Task.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Message{"failed to create task"})
		return
	}

	c.JSON(http.StatusOK, Message{fmt.Sprint(taskID)})
}

// DeleteTask deletes a task
// @Summary Delete a task
// @Description Delete a task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} Message "Task deleted"
// @Failure 400 {object} Message "ID should be an integer"
// @Failure 500 {object} Message "Failed to delete task"
// @Router /tasks/{id} [delete]
func (h *Handler) DeleteTask(c *gin.Context) {
	log := h.log.With(slog.String("op", "Handler.Delete"))

	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Warn("failed to parse task id", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, Message{"id should be integer"})
		return
	}

	err = h.service.Task.DeleteTask(uint(id64))
	if err != nil {
		log.Error("failed to delete task", slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, Message{"failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, Message{"task deleted"})

}

// FinishTask marks a task as finished
// @Summary Finish a task
// @Description Mark a task as finished by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} Message "Task ended"
// @Failure 400 {object} Message "ID should be an integer"
// @Failure 500 {object} Message "Failed to finish task"
// @Router /tasks/{id}/finish [post]
func (h *Handler) FinishTask(c *gin.Context) {
	log := h.log.With(slog.String("op", "Handler.FinishTask"))

	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Warn("failed to parse task id", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, Message{"id should be integer"})
		return
	}

	err = h.service.Task.FinishTask(uint(id64))
	if err != nil {
		log.Error("failed to finish task", slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, Message{"failed to finish task"})
		return
	}

	c.JSON(http.StatusOK, Message{"task ended"})
}

// StartPeriod starts a period for a task
// @Summary Start a period for a task
// @Description Start a period for a task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} Message "Period started"
// @Failure 400 {object} Message "ID should be an integer"
// @Failure 400 {object} Message "Failed to start period. Period not finished"
// @Failure 500 {object} Message "Failed to start"
// @Router /tasks/{id}/start [post]
func (h *Handler) StartPeriod(c *gin.Context) {
	log := h.log.With(slog.String("op", "Handler.StartPeriod"))

	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Warn("failed to parse id", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, Message{"id should be integer"})
		return
	}

	err = h.service.StartPeriod(uint(id64))
	if err != nil {
		if errors.Is(err, storage.ErrPeriodNotFinished) {
			log.Warn("Failed to start period. Period not finished", slog.Uint64("task_id", id64), slog.Any("err", err))
			c.JSON(http.StatusBadRequest, Message{err.Error()})
			return
		}
		log.Error("failed to start period", slog.Uint64("task_id", id64), slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, Message{"failed to start "})
		return
	}

	log.Info("Period started", slog.Uint64("task_id", id64))
	c.JSON(http.StatusOK, Message{"period started"})
}

// EndPeriod ends a period for a task
// @Summary End a period for a task
// @Description End a period for a task by ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} Message "Period ended"
// @Failure 400 {object} Message "ID should be an integer"
// @Failure 400 {object} Message "Failed to end period. Period not started"
// @Failure 500 {object} Message "Failed to end"
// @Router /tasks/{id}/end [post]
func (h *Handler) EndPeriod(c *gin.Context) {
	log := h.log.With(slog.String("op", "Handler.EndPeriod"))

	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Warn("failed to parse id", slog.Any("err", err))
		c.JSON(http.StatusBadRequest, Message{"id should be integer"})
		return
	}

	err = h.service.EndPeriod(uint(id64))
	if err != nil {
		if errors.Is(err, storage.ErrPeriodNotStarted) {
			log.Warn("Failed to end period. Period not started", slog.Uint64("task_id", id64), slog.Any("err", err))
			c.JSON(http.StatusBadRequest, Message{err.Error()})
			return
		}
		log.Error("failed to end period", slog.Uint64("task_id", id64), slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, Message{"failed to end"})
		return
	}

	log.Info("Period ended", slog.Uint64("task_id", id64))
	c.JSON(http.StatusOK, Message{"period ended"})
}
