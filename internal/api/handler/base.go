package handler

import (
	"logiflowCore/internal/config"
	"logiflowCore/pkg/response"
	"net/http"
)

func BaseHandler(w http.ResponseWriter, req *http.Request) {
	body := response.InitBody(nil, http.StatusOK, "Welcome to Logiflow Core")
	response.WrapRes(w, body)
}

func HealthHandler(w http.ResponseWriter, req *http.Request) {
	data := map[string]string{
		"status": "ok",
		"version": config.Version,
	}
	body := response.InitBody(data, http.StatusOK, "")
	response.WrapRes(w, body)
}