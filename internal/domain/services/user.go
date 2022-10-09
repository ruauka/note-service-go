package services

import (
	"web/internal/adapters/storage"
	"web/internal/domain/entities/dto"
)

// userService user service struct.
type userService struct {
	storage storage.UserStorage
}

// NewUserService user service func builder.
func NewUserService(userStorage storage.UserStorage) UserService {
	return &userService{storage: userStorage}
}

// GetUserByID get user by ID.
func (u *userService) GetUserByID(id string) (*dto.UserResp, error) {
	return u.storage.GetUserByID(id)
}

// GetAllUsers get all users.
func (u *userService) GetAllUsers() ([]dto.UserResp, error) {
	return u.storage.GetAllUsers()
}

// UpdateUser update user by ID.
func (u *userService) UpdateUser(newUser *dto.UserUpdate, userID string) error {
	return u.storage.UpdateUser(newUser, userID)
}

// DeleteUser delete user by ID.
func (u *userService) DeleteUser(id string) (int, error) {
	return u.storage.DeleteUser(id)
}
