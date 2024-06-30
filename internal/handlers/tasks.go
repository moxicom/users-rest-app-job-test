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

func (h *Handler) CreateTask(c *gin.Context) {
	log := h.log.With(slog.String("op", "Handler.CreateTask"))
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Error("error while parsing json ", err)
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

func (h *Handler) DeleteTask(c *gin.Context) {
	log := h.log.With(slog.String("op", "Handler.Delete"))

	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Warn("failed to parse task id", err)
		c.JSON(http.StatusBadRequest, Message{"id should be integer"})
		return
	}

	err = h.service.Task.DeleteTask(uint(id64))
	if err != nil {
		log.Error("failed to delete task", err)
		c.JSON(http.StatusInternalServerError, Message{"failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, Message{"task deleted"})

}

func (h *Handler) FinishTask(c *gin.Context) {
	log := h.log.With(slog.String("op", "Handler.FinishTask"))

	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Warn("failed to parse task id", err)
		c.JSON(http.StatusBadRequest, Message{"id should be integer"})
		return
	}

	err = h.service.Task.FinishTask(uint(id64))
	if err != nil {
		log.Error("failed to finish task", err)
		c.JSON(http.StatusInternalServerError, Message{"failed to finish task"})
		return
	}

	c.JSON(http.StatusOK, Message{"task ended"})
}

func (h *Handler) StartPeriod(c *gin.Context) {
	log := h.log.With(slog.String("op", "Handler.StartPeriod"))

	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Warn("failed to parse id", err)
		c.JSON(http.StatusBadRequest, Message{"id should be integer"})
		return
	}

	err = h.service.StartPeriod(uint(id64))
	if err != nil {
		if errors.Is(err, storage.ErrPeriodNotFinished) {
			log.Warn("Failed to start period. Period not finished", slog.Uint64("task_id", id64), err)
			c.JSON(http.StatusBadRequest, Message{err.Error()})
			return
		}
		log.Error("failed to start period", slog.Uint64("task_id", id64), err)
		c.JSON(http.StatusInternalServerError, Message{"failed to start "})
		return
	}

	log.Info("Period started", slog.Uint64("task_id", id64))
	c.JSON(http.StatusOK, Message{"period started"})
}

func (h *Handler) EndPeriod(c *gin.Context) {
	log := h.log.With(slog.String("op", "Handler.EndPeriod"))

	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Warn("failed to parse id", err)
		c.JSON(http.StatusBadRequest, Message{"id should be integer"})
		return
	}

	err = h.service.EndPeriod(uint(id64))
	if err != nil {
		if errors.Is(err, storage.ErrPeriodNotStarted) {
			log.Warn("Failed to end period. Period not started", slog.Uint64("task_id", id64), err)
			c.JSON(http.StatusBadRequest, Message{err.Error()})
			return
		}
		log.Error("failed to end period", slog.Uint64("task_id", id64), err)
		c.JSON(http.StatusInternalServerError, Message{"failed to end"})
		return
	}

	log.Info("Period ended", slog.Uint64("task_id", id64))
	c.JSON(http.StatusOK, Message{"period ended"})
}
