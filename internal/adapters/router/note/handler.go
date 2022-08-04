package note

import (
	"github.com/julienschmidt/httprouter"

	"web/internal/adapters/router/middleware"
	"web/internal/domain/services"
	"web/internal/utils"
)

type handler struct {
	service       services.Services
	logMiddleware func(next httprouter.Handle) httprouter.Handle
}

func NewHandler(service *services.Services, logFn func(next httprouter.Handle) httprouter.Handle) *handler {
	return &handler{
		service:       *service,
		logMiddleware: logFn,
	}
}

func Register(router *httprouter.Router, service *services.Services, logFn func(next httprouter.Handle) httprouter.Handle) {
	h := NewHandler(service, logFn)

	router.GET(utils.NoteURL, h.logMiddleware(middleware.CheckToken(h.GetNoteByID, h.service.Auth)))
	router.GET(utils.NotesURL, h.logMiddleware(middleware.CheckToken(h.GetAllNotesByUser, h.service.Auth)))
	router.POST(utils.NotesURL, h.logMiddleware(middleware.CheckToken(h.CreateNote, h.service.Auth)))
	router.PUT(utils.NoteURL, h.logMiddleware(middleware.CheckToken(h.UpdateNote, h.service.Auth)))
	router.DELETE(utils.NoteURL, h.logMiddleware(middleware.CheckToken(h.DeleteNote, h.service.Auth)))

	router.PUT(utils.TagsSet, h.logMiddleware(middleware.CheckToken(h.SetTags, h.service.Auth)))
	router.PUT(utils.TagsRemove, h.logMiddleware(middleware.CheckToken(h.RemoveTags, h.service.Auth)))
	router.GET(utils.AllTagsByNotes, h.logMiddleware(middleware.CheckToken(h.GetAllNotesWithTags, h.service.Auth)))
	router.GET(utils.AllTagsByNote, h.logMiddleware(middleware.CheckToken(h.GetNoteWithAllTags, h.service.Auth)))
}
