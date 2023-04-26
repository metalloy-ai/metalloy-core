package tools

import (
	"errors"
	"metalloyCore/pkg/response"
	"net/http"
)

var ErrInvalidCredentials error = errors.New("invalid credentials")
var ErrInvalidReqBody error = errors.New("invalid request body")
var ErrMissingParams error = errors.New("missing params")
var ErrUserNotFound error = errors.New("user not found")
var ErrUserAlreadyExist error = errors.New("user already exist")

func HandleError(err error, w http.ResponseWriter) bool {
	if errors.Is(err, ErrInvalidCredentials) {
		body := response.InitRes(http.StatusUnauthorized, "Unauthorized: login failed, invalid username or password", nil)
		response.WrapRes(w, body)
		return false
	} else if errors.Is(err, ErrInvalidReqBody) {
		body := response.InitRes(http.StatusBadRequest, "Bad request: Unable to process request body", nil)
		response.WrapRes(w, body)
		return false
	} else if errors.Is(err, ErrUserNotFound) {
		body := response.InitRes(http.StatusNotFound, "Not Found: User was not found", nil)
		response.WrapRes(w, body)
		return false
	} else if errors.Is(err, ErrUserAlreadyExist) {
		body := response.InitRes(http.StatusConflict, "Conflict: User already exist", nil)
		response.WrapRes(w, body)
		return false
	} else if errors.Is(err, ErrMissingParams) {
		body := response.InitRes(http.StatusBadRequest, "Bad request: Missing or Empty parameter", nil)
		response.WrapRes(w, body)
		return false
	} else if err != nil {
		body := response.InitRes(http.StatusInternalServerError, "Internal server error", nil)
		response.WrapRes(w, body)
		return false
	}
	return true
}