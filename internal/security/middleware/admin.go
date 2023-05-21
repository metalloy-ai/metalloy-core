package middleware

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

func AdminMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		return next(w, req)
	}
}