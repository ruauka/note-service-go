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

	//router.GET(utils.NotesURL, h.CheckToken)
	//router.GET(utils.UserURL, h.GetUserByID)
	router.POST(utils.NotesURL, middleware.CheckToken(h.CreateNote, h.service.Auth))
	//router.PUT(utils.UserURL, h.UpdateUser)
	//router.DELETE(utils.UserURL, h.DeleteUser)
}
