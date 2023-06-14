package tools

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"

	"metalloyCore/pkg/response"
)

type ErrInvalidCredentials struct{}
type ErrUserNotFound struct{}
type ErrUserAlreadyExist struct{}
type ErrFailedUsers struct {
	Users       map[string]interface{}
	FailedUsers []pgx.Row
}
type ErrForbiddenAccess struct{ Message string }
type ErrUnAuthorized struct{ Message string }
type ErrBadRequest struct{ Message string }

func (e ErrUserNotFound) Error() string     { return "user not found" }
func (e ErrUserAlreadyExist) Error() string { return "user already exists" }
func (e ErrFailedUsers) Error() string {
	return fmt.Sprintf("%d users failed to load", len(e.FailedUsers))
}
func (e ErrForbiddenAccess) Error() string { return e.Message }
func (e ErrUnAuthorized) Error() string    { return e.Message }
func (e ErrBadRequest) Error() string      { return e.Message }

func HandleError(err error, w http.ResponseWriter) bool {
	switch err := err.(type) {
	case ErrUserNotFound:
		body := response.InitRes(http.StatusNotFound, "Not Found: User was not found", nil)
		response.WrapRes(w, body)
	case ErrUserAlreadyExist:
		body := response.InitRes(http.StatusConflict, "Conflict: User already exists", nil)
		response.WrapRes(w, body)
	case ErrFailedUsers:
		body := response.InitRes(http.StatusInternalServerError, err.Error(), err.Users)
		response.WrapRes(w, body)
	case ErrUnAuthorized:
		body := response.InitRes(http.StatusUnauthorized, err.Error(), nil)
		response.WrapRes(w, body)
	case ErrBadRequest:
		body := response.InitRes(http.StatusBadRequest, err.Error(), nil)
		response.WrapRes(w, body)
	case ErrForbiddenAccess:
		body := response.InitRes(http.StatusForbidden, err.Error(), nil)
		response.WrapRes(w, body)
	case nil:
		return true
	default:
		errMsg := "Internal server error: " + err.Error()
		body := response.InitRes(http.StatusInternalServerError, errMsg, nil)
		response.WrapRes(w, body)
	}
	return false
}

func HandleEmptyError(input interface{}, err error) (interface{}, error) {
	if errors.Is(err, pgx.ErrNoRows) {
		return input, ErrUserNotFound{}
	}

	if err != nil {
		return nil, err
	}

	return input, nil
}

func NewUnAuthorizedErr(msg string) ErrUnAuthorized {
	return ErrUnAuthorized{Message: msg}
}

func NewBadRequestErr(msg string) ErrBadRequest {
	return ErrBadRequest{Message: msg}
}

func NewForbiddenAccessErr(msg string) ErrForbiddenAccess {
	return ErrForbiddenAccess{Message: msg}
}
