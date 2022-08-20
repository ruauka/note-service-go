package tag

import (
	"github.com/julienschmidt/httprouter"

	"web/internal/adapters/router/middleware"
	"web/internal/domain/services"
	"web/internal/utils"
)

// handler struct.
type handler struct {
	service       services.Services
	logMiddleware utils.LogMiddleware
}

// NewHandler func builder.
func NewHandler(service *services.Services, logFn utils.LogMiddleware) *handler {
	return &handler{
		service:       *service,
		logMiddleware: logFn,
	}
}

// Register register tag handlers.
func Register(router *httprouter.Router, service *services.Services, logFn utils.LogMiddleware) {
	h := NewHandler(service, logFn)

	router.POST(utils.TagsURL, h.logMiddleware(middleware.CheckToken(h.CreateTag, h.service.Auth)))
	router.GET(utils.TagURL, h.logMiddleware(middleware.CheckToken(h.GetTagByID, h.service.Auth)))
	router.GET(utils.TagsURL, h.logMiddleware(middleware.CheckToken(h.GetAllTagsByUser, h.service.Auth)))
	router.PUT(utils.TagURL, h.logMiddleware(middleware.CheckToken(h.UpdateTag, h.service.Auth)))
	router.DELETE(utils.TagURL, h.logMiddleware(middleware.CheckToken(h.DeleteTag, h.service.Auth)))
}
