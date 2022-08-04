package user

import (
	"github.com/julienschmidt/httprouter"

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

	router.POST(utils.Register, h.logger.LogMiddleware(h.RegisterUser))
	router.POST(utils.Login, h.GenerateToken)
	router.GET(utils.UsersURL, h.GetAllUsers)
	router.GET(utils.UserURL, h.GetUserByID)
	router.PUT(utils.UserURL, h.UpdateUser)
	router.DELETE(utils.UserURL, h.DeleteUser)
}
