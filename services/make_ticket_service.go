package services

import (
	"bzdev/database"
	"bzdev/models"
)

func MakeTicketService(
	username string, 
	password string,
	newTicket models.Ticket,
) error {
	db := database.GetDB()
	return database.AssignTicketToUser(db, username, newTicket)
}
