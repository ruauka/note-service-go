package tag

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

// Register register tag handlers.
func Register(router *httprouter.Router, service *services.Services, logFn dictionary.LogMiddleware) {
	h := NewHandler(service, logFn)

	router.POST(dictionary.TagsURL, h.logMiddleware(middleware.CheckToken(h.CreateTag, h.service.Auth)))
	router.GET(dictionary.TagURL, h.logMiddleware(middleware.CheckToken(h.GetTagByID, h.service.Auth)))
	router.GET(dictionary.TagsURL, h.logMiddleware(middleware.CheckToken(h.GetAllTagsByUser, h.service.Auth)))
	router.PUT(dictionary.TagURL, h.logMiddleware(middleware.CheckToken(h.UpdateTag, h.service.Auth)))
	router.DELETE(dictionary.TagURL, h.logMiddleware(middleware.CheckToken(h.DeleteTag, h.service.Auth)))
}
