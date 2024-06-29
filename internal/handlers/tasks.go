package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moxicom/user_test/internal/models"
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

func (h *Handler) EndTask(c *gin.Context) {
	log := h.log.With(slog.String("op", "Handler.EndTask"))

	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Error("failed to end task", err)
		c.JSON(http.StatusBadRequest, Message{"id should be integer"})
		return
	}

	err = h.service.Task.EndTask(uint(id64))
	if err != nil {
		log.Error("failed to end task", err)
		c.JSON(http.StatusInternalServerError, Message{"failed to end task"})
		return
	}

	c.JSON(http.StatusOK, Message{"task ended"})
}

func (h *Handler) DeleteTask(c *gin.Context) {
	log := h.log.With(slog.String("op", "Handler.Delete"))

	id := c.Param("id")
	id64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		log.Error("failed to delete task", err)
		c.JSON(http.StatusBadRequest, Message{"id should be integer"})
		return
	}

	//TODO service - delete task
	err = h.service.Task.DeleteTask(uint(id64))
	if err != nil {
		log.Error("failed to delete task", err)
		c.JSON(http.StatusInternalServerError, Message{"failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, Message{"task deleted"})

}
