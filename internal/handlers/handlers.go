package handlers

import (
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moxicom/user_test/internal/services"
)

type Message struct {
	Msg string `json:"msg"`
}

type Handler struct {
	service services.Service
	log     *slog.Logger
}

func New(service *services.Service, log *slog.Logger) *Handler {
	return &Handler{
		service: *service,
		log:     log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := router.Group("/users")
	{
		users.GET("/", h.GetUsers)
		users.POST("/", h.CreateUser)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
		users.GET("/:id/tasks", h.GetUsersWithTasks)
	}

	tasks := router.Group("/tasks")
	{
		tasks.POST("/", h.CreateTask)
		tasks.POST("/:id/end", h.EndTask)
		tasks.DELETE("/:id", h.DeleteTask)
	}

	return router
}
