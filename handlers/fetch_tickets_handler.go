package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bzdev/services"
	"bzdev/models"
)

func FetchTicketsHandler(c *gin.Context) {
	var req models.FetchTicketsRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
		})
		return
	}
	
	tickets, err := services.FetchTicketsService(req.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "failed to fetch tickets",
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"response": "fetch tickets successful",
			"tickets": tickets,
		})
	}
}