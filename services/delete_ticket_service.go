package services

import (
	"bzdev/database"
)

func DeleteTicketService(username string, title string) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}

	err = database.DeleteTicketByTitle(db, username, title)
	if err != nil {
		return err
	}

	return nil
}