package models

type User struct {
	ID       string
	Username string
	Password string
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
