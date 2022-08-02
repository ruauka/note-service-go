package services

import (
	"web/internal/adapters/storage"
	"web/internal/domain/enteties/dto"
)

type userService struct {
	storage storage.UserStorage
	// logger
}

func NewUserService(userStorage storage.UserStorage) UserService {
	return &userService{storage: userStorage}
}

func (u *userService) GetAllUsers() ([]dto.UserResp, error) {
	return u.storage.GetAllUsers()
}

func (u *userService) GetUserByID(id string) (*dto.UserResp, error) {
	return u.storage.GetUserByID(id)
}

func (u *userService) UpdateUser(newUser *dto.UserUpdate, userId string) error {
	return u.storage.UpdateUser(newUser, userId)
}

func (u *userService) DeleteUser(id string) (int, error) {
	return u.storage.DeleteUser(id)
}
