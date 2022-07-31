package services

import (
	"web/internal/adapters/storage"
	"web/internal/domain/enteties/dto"
	"web/internal/domain/errors"
)

type userService struct {
	storage storage.UserStorage
	// logger
}

func NewUserService(db storage.UserStorage) UserService {
	return &userService{storage: db}
}

func (u *userService) GetAllUsers() ([]dto.UserResp, error) {
	users, err := u.storage.GetAllUsers()
	if len(users) == 0 {
		return nil, errors.ErrUsersListEmpty
	}
	return users, err
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
