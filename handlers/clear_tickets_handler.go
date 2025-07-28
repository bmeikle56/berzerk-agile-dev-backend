package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bzdev/services"
	"bzdev/models"
)

func ClearTicketsHandler(c *gin.Context) {
	var req models.FetchTicketsRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
		})
		return
	}
	
	err := services.ClearTicketsService(req.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "failed to clear tickets",
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"response": "clear tickets successful",
		})
	}
}