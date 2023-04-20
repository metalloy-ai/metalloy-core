package response

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

type Response struct {
	Code 	int 		`json:"code"`
	Message string		`json:"message"`
	Data 	interface{} `json:"body"`
}

func InitRes(code int, message string, data interface{}) Response {
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
