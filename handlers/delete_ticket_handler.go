package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bzdev/services"
	"bzdev/models"
)

func DeleteTicketHandler(c *gin.Context) {
	var req models.DeleteTicketRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
		})
		return
	}
	
	err := services.DeleteTicketService(req.Username, req.Key)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "failed to delete ticket",
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"response": "delete ticket successful",
		})
	}
}