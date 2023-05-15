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

func (ac *AuthController) LoginHandler(w http.ResponseWriter, req *http.Request) {
	loginBody := auth.LoginRequest{}

	err := loginBody.DecodeBody(req.Body)
	if !tools.HandleError(err, w) {
		return
	}

	err = loginBody.Validate()
	if !tools.HandleError(err, w) {
		return
	}

	user, err := ac.Svc.Login(req.Context(), loginBody.Username, loginBody.Password)

	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "", nil)
	w.Header().Set("Authorization", "Bearer "+user.Token)
	response.WrapRes(w, &body)
}

func (ac *AuthController) RegisterHandler(w http.ResponseWriter, req *http.Request) {
	registerBody := &user.UserCreate{}

	err := registerBody.DecodeBody(req.Body)
	if !tools.HandleError(err, w) {
		return
	}

	err = registerBody.Validate()
	if !tools.HandleError(err, w) {
		return
	}

	user, err := ac.Svc.Register(req.Context(), registerBody)

	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusCreated, "", *user)
	response.WrapRes(w, &body)
}

func (ac *AuthController) ForgetPasswordHandler(w http.ResponseWriter, req *http.Request) {
	body := *response.InitRes(http.StatusOK, "", nil)
	response.WrapRes(w, &body)
}
