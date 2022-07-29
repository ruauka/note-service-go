package model

type User struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
