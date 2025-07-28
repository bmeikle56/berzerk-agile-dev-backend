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

// a status-less ticket
type NewTicket struct {
	Repo   string `json:"repo"`
	Tag    string `json:"tag"`
	Title  string `json:"title"`
	Dev    string `json:"dev"` // remove later...
	Notes  string `json:"notes"`
}

func (nt NewTicket) ToTicketWithStatus(status string) Ticket {
	return Ticket{
		Repo:   nt.Repo,
		Tag:    nt.Tag,
		Title:  nt.Title,
		Dev:    nt.Dev,
		Notes:  nt.Notes,
		Status: status,
	}
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewTicketRequest struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Ticket   NewTicket `json:"ticket"`
}

type FetchTicketsRequest struct {
	Username string `json:"username"`
}

type DeleteTicketRequest struct {
	Username string `json:"username"`
	Title    string `json:"title"`
}

type UpdateStatusRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Title    string `json:"title"`
	Status   string `json:"status"`
}

