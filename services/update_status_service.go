package services

import (
	"bzdev/database"
)

func UpdateStatusService(username string, repo string, key string, newStatus string) error {
	db := database.GetDB()
	err := database.UpdateTicketStatusByRepo(db, username, repo, key, newStatus)
	if err != nil {
		return err
	}
	return nil
}