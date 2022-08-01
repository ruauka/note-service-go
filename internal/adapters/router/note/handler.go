package note

import (
	"github.com/julienschmidt/httprouter"

	"web/internal/adapters/router/middleware"
	"web/internal/domain/services"
	"web/internal/utils"
)

type handler struct {
	service services.Services
	// logger
}

func NewHandler(service *services.Services) *handler {
	return &handler{
		service: *service,
	}
}

func Register(router *httprouter.Router, service *services.Services) {
	h := NewHandler(service)

	router.GET(utils.NoteURL, middleware.CheckToken(h.GetNoteByID, h.service.Auth))
	router.GET(utils.NotesURL, middleware.CheckToken(h.GetAllNotesByUser, h.service.Auth))
	router.POST(utils.NotesURL, middleware.CheckToken(h.CreateNote, h.service.Auth))
	router.PUT(utils.NoteURL, middleware.CheckToken(h.UpdateNote, h.service.Auth))
	router.DELETE(utils.NoteURL, middleware.CheckToken(h.DeleteNote, h.service.Auth))
}
