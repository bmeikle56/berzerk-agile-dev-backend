package handlers

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"bzdev/services"
	"bzdev/models"
)

func MakeTicketHandler(c *gin.Context) {
	var req models.TicketRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
		})
		return
	}
	
	err := services.MakeTicketService(req.Username, req.Password, req.Ticket)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "failed to make ticket",
			"error": err.Error(),
		})
	} else {
		authToken := os.Getenv("AUTH_TOKEN")
		c.JSON(http.StatusOK, gin.H{
			"response": "make ticket successful",
			"token": authToken,
		})
	}
}