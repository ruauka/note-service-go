package user

import (
	"github.com/julienschmidt/httprouter"

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

	router.POST(utils.Register, h.RegisterUser)
	router.POST(utils.Auth, h.GenerateToken)
	router.GET(utils.UsersURL, h.Identity(h.GetAllUsers))
	router.GET(utils.UserURL, h.GetUserByID)
	router.PUT(utils.UserURL, h.UpdateUser)
	router.DELETE(utils.UserURL, h.DeleteUser)
}
