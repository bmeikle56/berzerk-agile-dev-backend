package services

import (
	"bzdev/database"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func SignupService(username string, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}
	userExists, err := database.CheckIfUserExists(db, username)
	if err != nil {
		return err
	} else if userExists {
		return fmt.Errorf("user already exists")
	}
	err = database.InsertUser(db, username, string(hashed), 100)
	return err
}