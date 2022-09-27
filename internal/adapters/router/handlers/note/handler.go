// Package note Package note
package note

import (
	"github.com/julienschmidt/httprouter"

	"web/internal/adapters/router/middleware"
	"web/internal/domain/services"
	"web/internal/utils"
)

// Handler struct.
type Handler struct {
	service       services.Services
	logMiddleware utils.LogMiddleware
}

// NewHandler func builder.
func NewHandler(service *services.Services, logFn utils.LogMiddleware) *Handler {
	return &Handler{
		service:       *service,
		logMiddleware: logFn,
	}
}

// Register register note handlers.
func Register(router *httprouter.Router, service *services.Services, logFn utils.LogMiddleware) {
	h := NewHandler(service, logFn)

	router.POST(utils.NotesURL, h.logMiddleware(middleware.CheckToken(h.CreateNote, h.service.Auth)))
	router.GET(utils.NoteURL, h.logMiddleware(middleware.CheckToken(h.GetNoteByID, h.service.Auth)))
	router.GET(utils.NotesURL, h.logMiddleware(middleware.CheckToken(h.GetAllNotesByUser, h.service.Auth)))
	router.PUT(utils.NoteURL, h.logMiddleware(middleware.CheckToken(h.UpdateNote, h.service.Auth)))
	router.DELETE(utils.NoteURL, h.logMiddleware(middleware.CheckToken(h.DeleteNote, h.service.Auth)))

	router.PUT(utils.TagsSet, h.logMiddleware(middleware.CheckToken(h.SetTags, h.service.Auth)))
	router.PUT(utils.TagsRemove, h.logMiddleware(middleware.CheckToken(h.RemoveTags, h.service.Auth)))
	router.GET(utils.AllTagsByNotes, h.logMiddleware(middleware.CheckToken(h.GetAllNotesWithTags, h.service.Auth)))
	router.GET(utils.AllTagsByNote, h.logMiddleware(middleware.CheckToken(h.GetNoteWithAllTags, h.service.Auth)))
}
