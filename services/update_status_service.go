package services

import (
	"bzdev/database"
	// "fmt"
)

func UpdateStatusService(username string, title string, newStatus string) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	// for now assume the intended user is correct
	// userExists, err := database.CheckIfUserExists(db, username)
	// if err != nil {
	// 	return err
	// } else if userExists {
	// 	return fmt.Errorf("user already exists")
	// }
	database.UpdateTicketStatus(db, username, title, newStatus)

	return err
}