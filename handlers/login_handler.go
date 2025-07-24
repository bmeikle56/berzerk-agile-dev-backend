package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bzdev/services"
	"bzdev/models"
)

func LoginHandler(c *gin.Context) {
	var req models.AuthRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
		})
		return
	}
	
	err := services.LoginService(req.Username, req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "failed to login user",
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"response": "login successful",
		})
	}
}