// Package dto Package dto
package dto

// NoteUpdate dto.
type NoteUpdate struct {
	Title *string `json:"title"`
	Info  *string `json:"info"`
}

// NoteResp dto.
type NoteResp struct {
	Title string `json:"title"`
	Info  string `json:"info"`
}

// NotesResp dto.
type NotesResp struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Info  string `json:"info"`
}

// NoteWithTags dto.
type NoteWithTags struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Info     string `json:"info"`
	TagsResp `json:"tags"`
}

// NoteWithTagsResp dto.
type NoteWithTagsResp struct {
	ID       string     `json:"id"`
	Title    string     `json:"title"`
	Info     string     `json:"info"`
	TagsResp []TagsResp `json:"tags"`
}
