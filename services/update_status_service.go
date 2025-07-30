package services

import (
	"bzdev/database"
	// "fmt"
)

func UpdateStatusService(username string, title string, newStatus string) error {
	db := database.GetDB()
	// for now assume the intended user is correct
	// userExists, err := database.CheckIfUserExists(db, username)
	// if err != nil {
	// 	return err
	// } else if userExists {
	// 	return fmt.Errorf("user already exists")
	// }
	err := database.UpdateTicketStatus(db, username, title, newStatus)
	if err != nil {
		return err
	}

	return nil
}