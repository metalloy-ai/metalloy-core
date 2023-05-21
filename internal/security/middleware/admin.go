package middleware

import (
	"metalloyCore/pkg/validator"
	"metalloyCore/tools"
	"net/http"
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