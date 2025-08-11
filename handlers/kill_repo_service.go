package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bzdev/services"
	"bzdev/models"
)

func KillRepoHandler(c *gin.Context) {
	var req models.KillRepoRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
		})
		return
	}
	
	err := services.KillRepoService(req.Username, req.Repo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "failed to kill repo",
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"response": "repo killed successful",
		})
	}
}