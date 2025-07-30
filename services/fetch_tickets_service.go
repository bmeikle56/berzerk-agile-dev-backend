package services

import (
	"bzdev/database"
	"bzdev/models"
)

func FetchTicketsService(username string) (models.UserData, error) {
	db := database.GetDB()

	tickets, err := database.FetchUserData(db, username)
	if err != nil {
		return models.UserData{}, err
	}

	return tickets, nil
}