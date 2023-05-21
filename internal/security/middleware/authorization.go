package middleware

import (
	"context"
	"net/http"

	"github.com/uptrace/bunrouter"

	"metalloyCore/internal/config"
	"metalloyCore/internal/security/jwt"
	"metalloyCore/tools"
)

type contextKey string
const CtxPayloadKey contextKey = "userPayload"

func Authorization(cfg config.Setting) bunrouter.MiddlewareFunc {
	jwtHandler := jwt.InitJWTHandler(cfg)
	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		return func(w http.ResponseWriter, req bunrouter.Request) error {
			claims, err := jwtHandler.ValidateRequest(req.Request)
			if err != nil {
				tools.HandleError(err, w)
				return nil
			}

			ctx := context.WithValue(req.Context(), CtxPayloadKey, &claims.UserPayload)
			req = req.WithContext(ctx)

			return next(w, req)
		}
	}
}
