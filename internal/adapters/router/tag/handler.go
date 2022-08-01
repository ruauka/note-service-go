package tag

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

	router.GET(utils.TagURL, middleware.CheckToken(h.GetTagByID, h.service.Auth))
	router.GET(utils.TagsURL, middleware.CheckToken(h.GetAllTagsByUser, h.service.Auth))
	router.POST(utils.TagsURL, middleware.CheckToken(h.CreateTag, h.service.Auth))
	router.PUT(utils.TagURL, middleware.CheckToken(h.UpdateTag, h.service.Auth))
	router.DELETE(utils.TagURL, middleware.CheckToken(h.DeleteTag, h.service.Auth))
}
