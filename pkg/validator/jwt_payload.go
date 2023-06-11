package validator

import (
	"net/http"

	"metalloyCore/internal/security/jwt"
	"metalloyCore/tools"
)

func ValidatePayload(req *http.Request, username string) error {
	payload := req.Context().Value(tools.CtxPayloadKey).(*jwt.UserPayload)

	if payload.Username != username {
		return tools.NewForbiddenAccessErr("You are not allowed to access this resource")
	}

	return nil
}

func ValidatePayloadAdmin(req *http.Request) error {
	payload := req.Context().Value(tools.CtxPayloadKey).(*jwt.UserPayload)

	if payload.Role != "admin" {
		return tools.NewForbiddenAccessErr("You are not an admin")
	}

	return nil
}
