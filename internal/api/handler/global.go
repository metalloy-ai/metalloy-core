package handler

import (
	"net/http"

	"metalloyCore/pkg/response"
)

func EmptyParamHandler(w http.ResponseWriter, req *http.Request) {
	body := *response.InitRes(http.StatusBadRequest, "Bad request: empty parameter", nil)
	response.WrapRes(w, &body)
}
