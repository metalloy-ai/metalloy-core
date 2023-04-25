package handler

import (
	"metalloyCore/internal/security/auth"
	"net/http"
)

type AuthController struct {
	Svc auth.AuthService
}

func InitAuthController(svc auth.AuthService) *AuthController {
	return &AuthController{Svc: svc}
}

func (ac AuthController) LoginHandler(w http.ResponseWriter, req *http.Request) {

}

func (ac AuthController) RegisterHandler(w http.ResponseWriter, req *http.Request) {

}
