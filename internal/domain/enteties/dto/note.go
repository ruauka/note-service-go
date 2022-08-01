package dto

type NoteUpdate struct {
	Title *string `json:"title"`
	Info  *string `json:"info"`
}

type NoteResp struct {
	Title string `json:"title"`
	Info  string `json:"info"`
}

type NotesResp struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Info  string `json:"info"`
}
