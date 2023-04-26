package handler

import (
	"net/http"

	"metalloyCore/internal/security/auth"
	"metalloyCore/pkg/response"
	"metalloyCore/tools"
)

type AuthController struct {
	Svc auth.AuthService
}

func InitAuthController(svc auth.AuthService) *AuthController {
	return &AuthController{Svc: svc}
}

func (ac AuthController) LoginHandler(w http.ResponseWriter, req *http.Request) {
	loginBody := auth.LoginReq{}
	
	if !tools.HandleError(loginBody.DecodeBody(req.Body), w) {
		return
	}
	if !tools.HandleError(loginBody.Validate(), w) {
		return
	}

	user, err := ac.Svc.Login(loginBody.Username, loginBody.Password)

	if !tools.HandleError(err, w) {
		return
	}

	body := response.InitRes(http.StatusOK, "", user)
	response.WrapRes(w, body)
}

func (ac AuthController) RegisterHandler(w http.ResponseWriter, req *http.Request) {}
