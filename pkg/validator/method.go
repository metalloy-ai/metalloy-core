package validator

import (
	"logiflowCore/pkg/response"
	"net/http"
)

func ValidateMethodGet(w http.ResponseWriter, req *http.Request) bool {
	if req.Method != http.MethodGet {
		body := response.InitRes(http.StatusMethodNotAllowed, "Method not allowed", nil)
		response.WrapRes(w, body)
		return false
	}
	return true
}