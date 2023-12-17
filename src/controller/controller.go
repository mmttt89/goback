package controller

import (
	"goback/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context, userService *service.UserService, userID string) {
	user, err := userService.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
