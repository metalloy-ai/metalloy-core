package middleware

import (
	"net/http"

	"metalloyCore/pkg/validator"
	"metalloyCore/tools"
)

func Admin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := validator.ValidatePayloadAdmin(r)
		if !tools.HandleError(err, w) {
			return
		}
		next(w, r)
	}
}
