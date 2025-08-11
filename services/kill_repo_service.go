package services

import (
	"bzdev/database"
)

func KillRepoService(username string, repo string) error {
	db := database.GetDB()

	err := database.KillRepoByName(db, username, repo)
	if err != nil {
		return err
	}

	return nil
}