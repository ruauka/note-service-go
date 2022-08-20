// Package user Package user
package user

import (
	"github.com/julienschmidt/httprouter"

	"web/internal/domain/services"
	"web/internal/utils"
)

// handler struct.
type handler struct {
	service       services.Services
	LogMiddleware utils.LogMiddleware
}

// NewHandler func builder.
func NewHandler(service *services.Services, logFn utils.LogMiddleware) *handler {
	return &handler{
		service:       *service,
		LogMiddleware: logFn,
	}
}

// Register register user handlers.
func Register(router *httprouter.Router, service *services.Services, logFn utils.LogMiddleware) {
	h := NewHandler(service, logFn)

	router.POST(utils.Register, h.LogMiddleware(h.RegisterUser))
	router.POST(utils.Login, h.LogMiddleware(h.GenerateToken))
	router.GET(utils.UserURL, h.LogMiddleware(h.GetUserByID))
	router.GET(utils.UsersURL, h.LogMiddleware(h.GetAllUsers))
	router.PUT(utils.UserURL, h.LogMiddleware(h.UpdateUser))
	router.DELETE(utils.UserURL, h.LogMiddleware(h.DeleteUser))
}
