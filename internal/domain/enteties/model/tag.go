package model

type Tag struct {
	ID      string `json:"id"`
	TagName string `json:"tagname" validate:"required"`
	UserID  string `json:"user_id"`
}
