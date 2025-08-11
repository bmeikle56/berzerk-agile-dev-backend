package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bzdev/services"
	"bzdev/models"
)

func UpdateStatusHandler(c *gin.Context) {
	var req models.UpdateStatusRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
		})
		return
	}
	
	err := services.UpdateStatusService(req.Username, req.Repo, req.Title, req.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "failed to update status",
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"response": "update status successful",
		})
	}
}