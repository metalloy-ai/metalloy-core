package middleware

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

func CorsMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	allowedOrgins := []string{"http://localhost:3000"}

	return func(w http.ResponseWriter, req bunrouter.Request) error {
		origin := req.Header.Get("Origin")
		validOrigin := false

		for _, allowedOrgin := range allowedOrgins {
			if allowedOrgin == origin {
				validOrigin = true
				break
			}
		}
		
		if !validOrigin {
			return next(w, req)
		}

		h := w.Header()
		h.Set("Access-Control-Allow-Origin", origin)
		h.Set("Access-Control-Allow-Credentials", "true")

		if req.Method == http.MethodOptions {
			h.Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,HEAD")
			h.Set("Access-Control-Allow-Headers", "authorization,content-type")
			return nil
		}

		return next(w, req)
	}
}