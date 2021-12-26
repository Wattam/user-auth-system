package user_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-auth-system/database"
	"github.com/wattam/golang-auth-system/models"
)

func Get(c *gin.Context) {

	user := models.User{}
	id := c.Param("id")

	err := database.Db.First(&user, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
