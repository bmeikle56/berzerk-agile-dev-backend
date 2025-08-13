package services

import (
	"bzdev/database"
)

func DeleteTicketService(username string, key string) error {
	db := database.GetDB()

	err := database.DeleteTicketByKey(db, username, key)
	if err != nil {
		return err
	}

	return nil
}