package handler

import (
	"net/http"
	"strconv"

	"metalloyCore/internal/config"
	"metalloyCore/pkg/response"
)

func DefaultHandler(w http.ResponseWriter, req *http.Request) {
	body := *response.InitRes(http.StatusOK, "Welcome to Metalloy AI", nil)
	response.WrapRes(w, &body)
}

func HealthHandler(cfg config.Setting) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		data := map[string]string{
			"status":  "ok",
			"version": cfg.Version,
			"host":    cfg.Host,
			"port":    strconv.Itoa(cfg.Port),
		}
		body := *response.InitRes(http.StatusOK, "", data)
		response.WrapRes(w, &body)
	}
}
