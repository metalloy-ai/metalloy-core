package tools

import (
	"net/http"
	
	"metalloyCore/pkg/response"
)

type ErrInvalidCredentials struct{}
type ErrInvalidReqBody struct{}
type ErrMissingParams struct{}
type ErrUserNotFound struct{}
type ErrUserAlreadyExist struct{}

func (e ErrInvalidCredentials) Error() string {return "invalid credentials"}
func (e ErrInvalidReqBody) Error() string {return "invalid request body"}
func (e ErrMissingParams) Error() string {return "missing params"}
func (e ErrUserNotFound) Error() string {return "user not found"}
func (e ErrUserAlreadyExist) Error() string {return "user already exists"}

func HandleError(err error, w http.ResponseWriter) bool {
	switch err.(type) {
	case ErrInvalidCredentials:
		body := response.InitRes(http.StatusUnauthorized, "Unauthorized: login failed, invalid username or password", nil)
		response.WrapRes(w, body)
	case ErrInvalidReqBody:
		body := response.InitRes(http.StatusBadRequest, "Bad request: Unable to process request body", nil)
		response.WrapRes(w, body)
	case ErrUserNotFound:
		body := response.InitRes(http.StatusNotFound, "Not Found: User was not found", nil)
		response.WrapRes(w, body)
	case ErrUserAlreadyExist:
		body := response.InitRes(http.StatusConflict, "Conflict: User already exists", nil)
		response.WrapRes(w, body)
	case ErrMissingParams:
		body := response.InitRes(http.StatusBadRequest, "Bad request: Missing or Empty parameter", nil)
		response.WrapRes(w, body)
	case nil:
		return true
	default:
		body := response.InitRes(http.StatusInternalServerError, "Internal server error", nil)
		response.WrapRes(w, body)
	}
	return false
}