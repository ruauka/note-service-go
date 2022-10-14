// Package user Package user
package user

import (
	"github.com/julienschmidt/httprouter"

	"web/internal/domain/services"
	"web/internal/utils/dictionary"
)

// Handler struct.
type Handler struct {
	service       services.Services
	LogMiddleware dictionary.LogMiddleware
}

// NewHandler func builder.
func NewHandler(service *services.Services, logFn dictionary.LogMiddleware) *Handler {
	return &Handler{
		service:       *service,
		LogMiddleware: logFn,
	}
}

// Register register user handlers.
func Register(router *httprouter.Router, service *services.Services, logFn dictionary.LogMiddleware) {
	h := NewHandler(service, logFn)

	router.POST(dictionary.Register, h.LogMiddleware(h.RegisterUser))
	router.POST(dictionary.Login, h.LogMiddleware(h.GenerateToken))
	router.GET(dictionary.UserURL, h.LogMiddleware(h.GetUserByID))
	router.GET(dictionary.UsersURL, h.LogMiddleware(h.GetAllUsers))
	router.PUT(dictionary.UserURL, h.LogMiddleware(h.UpdateUser))
	router.DELETE(dictionary.UserURL, h.LogMiddleware(h.DeleteUser))
}
