package services

import (
	"bzdev/database"
)

func UpdateStatusService(username string, repo string, title string, newStatus string) error {
	db := database.GetDB()
	err := database.UpdateTicketStatusByRepo(db, username, repo, title, newStatus)
	if err != nil {
		return err
	}
	return nil
}