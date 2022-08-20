package dto

// UserUpdate dto.
type UserUpdate struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

// UserAuth dto.
type UserAuth struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserResp dto.
type UserResp struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
