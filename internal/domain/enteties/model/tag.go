package model

type Tag struct {
	ID      string `json:"id"`
	TagName string `json:"tagname"`
	UserID  string `json:"user_id"`
}
