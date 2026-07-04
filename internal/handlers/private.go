package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	userID, _ := c.Get("userID")

	c.JSON(http.StatusOK, gin.H{
		"userID":  userID,
		"message": "This is your profile.",
	})
}

func Settings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Settings page.",
	})
}
