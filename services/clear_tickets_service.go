package services

import (
	"bzdev/database"
)

func ClearTicketsService(username string) error {
	db := database.GetDB()

	err := database.DeleteAllTickets(db, username)
	if err != nil {
		return err
	}

	return nil
}