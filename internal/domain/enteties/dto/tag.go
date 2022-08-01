package dto

type TagResp struct {
	TagName string `json:"tagname"`
}

type TagUpdate struct {
	TagName *string `json:"tagname"`
}

type TagsResp struct {
	ID      string `json:"id"`
	TagName string `json:"tagname"`
}
