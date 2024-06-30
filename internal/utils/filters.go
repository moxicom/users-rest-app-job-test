package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/moxicom/user_test/internal/models"
)

func GetFilters(c *gin.Context) models.UserFilters {
	f := models.UserFilters{
		PassportNumber: c.Query("passport_number"),
		Surname:        c.Query("surname"),
		Name:           c.Query("name"),
		Patronymic:     c.Query("patronymic"),
		Address:        c.Query("address"),
	}
	return f
}
