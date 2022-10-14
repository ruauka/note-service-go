// Package note Package note
package note

import (
	"github.com/julienschmidt/httprouter"

	"web/internal/adapters/router/middleware"
	"web/internal/domain/services"
	"web/internal/utils/dictionary"
)

// Handler struct.
type Handler struct {
	service       services.Services
	logMiddleware dictionary.LogMiddleware
}

// NewHandler func builder.
func NewHandler(service *services.Services, logFn dictionary.LogMiddleware) *Handler {
	return &Handler{
		service:       *service,
		logMiddleware: logFn,
	}
}

// Register register note handlers.
func Register(router *httprouter.Router, service *services.Services, logFn dictionary.LogMiddleware) {
	h := NewHandler(service, logFn)

	router.POST(dictionary.NotesURL, h.logMiddleware(middleware.CheckToken(h.CreateNote, h.service.Auth)))
	router.GET(dictionary.NoteURL, h.logMiddleware(middleware.CheckToken(h.GetNoteByID, h.service.Auth)))
	router.GET(dictionary.NotesURL, h.logMiddleware(middleware.CheckToken(h.GetAllNotesByUser, h.service.Auth)))
	router.PUT(dictionary.NoteURL, h.logMiddleware(middleware.CheckToken(h.UpdateNote, h.service.Auth)))
	router.DELETE(dictionary.NoteURL, h.logMiddleware(middleware.CheckToken(h.DeleteNote, h.service.Auth)))

	router.PUT(dictionary.TagsSet, h.logMiddleware(middleware.CheckToken(h.SetTags, h.service.Auth)))
	router.PUT(dictionary.TagsRemove, h.logMiddleware(middleware.CheckToken(h.RemoveTags, h.service.Auth)))
	router.GET(dictionary.AllTagsByNotes, h.logMiddleware(middleware.CheckToken(h.GetAllNotesWithTags, h.service.Auth)))
	router.GET(dictionary.AllTagsByNote, h.logMiddleware(middleware.CheckToken(h.GetNoteWithAllTags, h.service.Auth)))
}
