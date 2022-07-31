package model

type Note struct {
	ID     string `json:"id" db:"id"`
	Title  string `json:"title"`
	Info   string `json:"info"`
	UserID string `json:"user_id"`
}
