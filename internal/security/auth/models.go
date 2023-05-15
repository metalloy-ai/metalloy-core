package auth

import "metalloyCore/internal/domain/user"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	*user.UserResponse
	Token string `json:"token"`
}
