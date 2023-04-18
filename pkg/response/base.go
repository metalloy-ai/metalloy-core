package response

import (
	"net/http"
	// "strconv"

	"github.com/uptrace/bunrouter"
)

type Response struct {
	Code 	int 		`json:"code"`
	Data 	interface{} `json:"body"`
	Message string		`json:"message"`
}

func InitBody(data interface{}, code int, message string) Response {
	return Response{
		Code: code,
		Message: message,
		Data: data,
	}
}

func WrapRes(w http.ResponseWriter, body Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(body.Code)
	_ = bunrouter.JSON(w, body)
}
