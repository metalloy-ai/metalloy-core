package handler

import (
	"net/http"

	"metalloyCore/internal/domain/user"
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

	err := loginBody.DecodeBody(req.Body)
	if !tools.HandleError(err, w) {
		return
	}
	
	err = loginBody.Validate()
	if !tools.HandleError(err, w) {
		return
	}

	user, err := ac.Svc.Login(loginBody.Username, loginBody.Password)

	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "", user)
	response.WrapRes(w, &body)
}

func (ac AuthController) RegisterHandler(w http.ResponseWriter, req *http.Request) {
	registerBody := user.UserCreate{}
	
	err := registerBody.DecodeBody(req.Body)
	if !tools.HandleError(err, w) {
		return
	}

	err = registerBody.Validate()
	if !tools.HandleError(err, w) {
		return
	}

	user, err := ac.Svc.Register(registerBody)

	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "", user)
	response.WrapRes(w, &body)
}
