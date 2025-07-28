package services

import (
	"bzdev/database"
)

func ClearTicketsService(username string) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}

	err = database.DeleteAllTickets(db, username)
	if err != nil {
		return err
	}

	return nil
}