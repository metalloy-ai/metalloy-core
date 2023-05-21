package validator

import (
	"net/http"

	"metalloyCore/internal/security/jwt"
	"metalloyCore/internal/security/middleware"
	"metalloyCore/tools"
)

func ValidatePayload(req *http.Request, username string) error {
	payload := req.Context().Value(middleware.CtxPayloadKey).(*jwt.UserPayload)

	if payload.Username != username {
		return tools.ErrForbiddenAccess{}
	}

	return nil
}
