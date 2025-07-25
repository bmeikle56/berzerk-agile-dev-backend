package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bzdev/services"
	"bzdev/models"
)

func MakeTicketHandler(c *gin.Context) {
	var req models.NewTicketRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
		})
		return
	}

	// convert NewTicket to Ticket and add default status
	ticket := req.Ticket.ToTicketWithStatus("new")
	
	err := services.MakeTicketService(req.Username, req.Password, ticket)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "failed to make ticket",
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"response": "make ticket successful",
		})
	}
}