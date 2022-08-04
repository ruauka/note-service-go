package user

import (
	"github.com/julienschmidt/httprouter"

	"web/internal/domain/services"
	"web/internal/utils"
)

type handler struct {
	service       services.Services
	logMiddleware utils.LogMiddleware
}

func NewHandler(service *services.Services, logFn utils.LogMiddleware) *handler {
	return &handler{
		service:       *service,
		logMiddleware: logFn,
	}
}

func Register(router *httprouter.Router, service *services.Services, logFn utils.LogMiddleware) {
	h := NewHandler(service, logFn)

	router.POST(utils.Register, h.logMiddleware(h.RegisterUser))
	router.POST(utils.Login, h.logMiddleware(h.GenerateToken))
	router.GET(utils.UsersURL, h.logMiddleware(h.GetAllUsers))
	router.GET(utils.UserURL, h.logMiddleware(h.GetUserByID))
	router.PUT(utils.UserURL, h.logMiddleware(h.UpdateUser))
	router.DELETE(utils.UserURL, h.logMiddleware(h.DeleteUser))
}
