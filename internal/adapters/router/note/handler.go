package note

import (
	"github.com/julienschmidt/httprouter"

	"web/internal/adapters/router/middleware"
	"web/internal/domain/services"
	"web/internal/utils"
	"web/pkg/logger"
)

type handler struct {
	service services.Services
	logger  logger.Logger
}

func NewHandler(service *services.Services, logger *logger.Logger) *handler {
	return &handler{
		service: *service,
		logger:  *logger,
	}
}

func Register(router *httprouter.Router, service *services.Services, logger *logger.Logger) {
	h := NewHandler(service, logger)

	router.GET(utils.NoteURL, middleware.CheckToken(h.GetNoteByID, h.service.Auth))
	router.GET(utils.NotesURL, middleware.CheckToken(h.GetAllNotesByUser, h.service.Auth))
	router.POST(utils.NotesURL, middleware.CheckToken(h.CreateNote, h.service.Auth))
	router.PUT(utils.NoteURL, middleware.CheckToken(h.UpdateNote, h.service.Auth))
	router.DELETE(utils.NoteURL, middleware.CheckToken(h.DeleteNote, h.service.Auth))

	router.PUT(utils.TagsSet, middleware.CheckToken(h.SetTags, h.service.Auth))
	router.PUT(utils.TagsRemove, middleware.CheckToken(h.RemoveTags, h.service.Auth))
	router.GET(utils.AllTagsByNotes, middleware.CheckToken(h.GetAllNotesWithTags, h.service.Auth))
	router.GET(utils.AllTagsByNote, middleware.CheckToken(h.GetNoteWithAllTags, h.service.Auth))
}
