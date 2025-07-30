package services

import (
	"bzdev/database"
)

func DeleteTicketService(username string, title string) error {
	db := database.GetDB()

	err := database.DeleteTicketByTitle(db, username, title)
	if err != nil {
		return err
	}

	return nil
}