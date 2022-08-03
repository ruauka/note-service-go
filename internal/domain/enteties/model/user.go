package model

type User struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
