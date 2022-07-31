package model

type Note struct {
	ID     string `json:"id" db:"id"`
	Note   string `json:"note"`
	UserID string `json:"user_id"`
}
