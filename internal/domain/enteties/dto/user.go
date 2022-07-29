package dto

type UserUpdate struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResp struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
