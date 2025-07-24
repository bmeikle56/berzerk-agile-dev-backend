package models

type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Data     UserData `json:"data"`
}

type UserData struct {
	Tickets []Ticket `json:"tickets"`
}

type Ticket struct {
	Repo   string `json:"repo"`
	Tag    string `json:"tag"`
	Title  string `json:"title"`
	Dev    string `json:"dev"`
	Notes  string `json:"notes"`
	Status string `json:"status"`
}

// I might not need this
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TicketRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Ticket   Ticket `json:"ticket"`
}
