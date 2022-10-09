package dto

// TagResp dto.
type TagResp struct {
	TagName string `json:"tagname"`
}

// TagUpdate dto.
type TagUpdate struct {
	TagName *string `json:"tagname"`
}

// TagsResp dto.
type TagsResp struct {
	ID      string `json:"id"`
	TagName string `json:"tagname"`
}
