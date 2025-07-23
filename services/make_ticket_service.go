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
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	return database.AssignTicketToUser(db, username, newTicket)
}
