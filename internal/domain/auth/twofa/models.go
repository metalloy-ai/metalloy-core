package twofa

import (
	"metalloyCore/internal/security/jwt"
)

type TwofaRequest struct {
	jwt.UserPayload
	Email string `json:"email"`
	Data  []byte `json:"data"`
}

type TwofaVerifyRequest struct {
	Username string `json:"username"`
	Code     int    `json:"code"`
}

type TwofaResponse struct {
	jwt.UserPayload
	Code int    `json:"code"`
	Data string `json:"data"`
}
