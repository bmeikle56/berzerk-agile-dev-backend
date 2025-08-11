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
	
	err := services.DeleteTicketService(req.Username, req.Title)

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

/**
Only one active ticket per repo
{
  bz-dev {
	  ticket-id  (green if active)
		ticket-id2
		ticket-id3
	}
	bz-dev-backend {
	  ticket-id  (green if active)
		ticket-id2
		ticket-id3
	}
	braeden-meikle-site {
	  ticket-id  (green if active)
		ticket-id2
		ticket-id3
	}
}
*/