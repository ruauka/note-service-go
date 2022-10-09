// Package model Package model
package model

// Note model.
type Note struct {
	ID     string `json:"id" db:"id"`
	Title  string `json:"title" validate:"required"`
	Info   string `json:"info" validate:"required"`
	UserID string `json:"user_id"`
}
