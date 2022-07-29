package user

import (
	"github.com/julienschmidt/httprouter"

	"web/internal/domain/services"
)

const (
	usersURL = "/users"
	userURL  = "/users/:id"
	register = "/register"
	auth     = "/auth"
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

	router.POST(register, h.RegisterUser)
	router.POST(auth, h.GenerateToken)
	router.GET(usersURL, h.GetAllUsers)
	router.GET(userURL, h.GetUserByID)
	router.PUT(userURL, h.UpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}
