package auth

import (
	"github.com/google/uuid"

	"metalloyCore/internal/domain/user"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type AuthPayload struct {
	UserID   uuid.UUID     `json:"user_id"`
	Username string        `json:"username"`
	UserType user.UserType `json:"user_type"`
}

type AuthGenRequest struct {
	AuthPayload
	Email string `json:"email"`
}
