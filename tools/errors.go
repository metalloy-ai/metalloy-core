package tools

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"

	"metalloyCore/pkg/response"
)

type Response interface{}
type ErrInvalidCredentials struct{}
type ErrInvalidReq struct{}
type ErrMissingParams struct{}
type ErrUserNotFound struct{}
type ErrUserAlreadyExist struct{}
type ErrParseClaims struct{}
type ErrExpiredToken struct{}
type ErrFailedUsers struct {
	Users       map[string]interface{}
	FailedUsers []pgx.Row
}

func (e ErrInvalidCredentials) Error() string { return "invalid credentials" }
func (e ErrInvalidReq) Error() string         { return "invalid request" }
func (e ErrMissingParams) Error() string      { return "missing params" }
func (e ErrUserNotFound) Error() string       { return "user not found" }
func (e ErrUserAlreadyExist) Error() string   { return "user already exists" }
func (e ErrFailedUsers) Error() string {
	return fmt.Sprintf("%d users failed to load", len(e.FailedUsers))
}
func (e ErrParseClaims) Error() string  { return "failed to parse claims" }
func (e ErrExpiredToken) Error() string { return "token has expired" }

func HandleError(err error, w http.ResponseWriter) bool {
	switch err := err.(type) {
	case ErrInvalidCredentials:
		body := response.InitRes(http.StatusUnauthorized, "Unauthorized: login failed, invalid username or password", nil)
		response.WrapRes(w, body)
	case ErrInvalidReq:
		body := response.InitRes(http.StatusBadRequest, "Bad request: Unable to process request due to the param or body", nil)
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
	case ErrFailedUsers:
		body := response.InitRes(http.StatusInternalServerError, err.Error(), err.Users)
		response.WrapRes(w, body)
	case ErrParseClaims:
		body := response.InitRes(http.StatusBadRequest, "Internal server error: Failed to parse claims", nil)
		response.WrapRes(w, body)
	case ErrExpiredToken:
		body := response.InitRes(http.StatusUnauthorized, "Unauthorized: Token has expired", nil)
		response.WrapRes(w, body)
	case nil:
		return true
	default:
		body := response.InitRes(http.StatusInternalServerError, "Internal server error", nil)
		response.WrapRes(w, body)
	}
	return false
}

func HandleEmptyError(input Response, err error) (interface{}, error) {
	if errors.Is(err, pgx.ErrNoRows) {
		return input, ErrUserNotFound{}
	}

	if err != nil {
		return nil, err
	}

	return input, nil
}
