package handler

import (
	"net/http"

	"metalloyCore/internal/domain/auth"
	"metalloyCore/internal/domain/auth/twofa"
	"metalloyCore/internal/domain/user"
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

	err = ac.Svc.Login(req.Context(), loginBody.Username, loginBody.Password)
	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "", nil)
	response.WrapRes(w, &body)
}

func (ac *AuthController) LoginVerifyHandler(w http.ResponseWriter, req *http.Request) {
	twofaReq := twofa.TwofaVerifyRequest{}

	err := twofaReq.DecodeBody(req.Body)
	if !tools.HandleError(err, w) {
		return
	}

	err = twofaReq.Validate()
	if !tools.HandleError(err, w) {
		return
	}

	auth, err := ac.Svc.LoginVerify(req.Context(), twofaReq.Username, twofaReq.Code)
	if !tools.HandleError(err, w) {
		return
	}

	response.InitAuthRes(w, http.StatusOK, auth.Token)
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

	err = ac.Svc.Register(req.Context(), registerBody)
	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "", nil)
	response.WrapRes(w, &body)
}

func (ac *AuthController) RegisterVerifyHandler(w http.ResponseWriter, req *http.Request) {
	twofaReq := twofa.TwofaVerifyRequest{}

	err := twofaReq.DecodeBody(req.Body)
	if !tools.HandleError(err, w) {
		return
	}

	err = twofaReq.Validate()
	if !tools.HandleError(err, w) {
		return
	}

	auth, err := ac.Svc.RegisterVerify(req.Context(), twofaReq.Username, twofaReq.Code)
	if !tools.HandleError(err, w) {
		return
	}

	response.InitAuthRes(w, http.StatusCreated, auth.Token)
}

func (ac *AuthController) ResetPasswordHandler(w http.ResponseWriter, req *http.Request) {
	email := req.URL.Query().Get("email")
	
	if email == "" {
		err := tools.NewBadRequestErr("email is required")
		tools.HandleError(err, w)
		return
	}

	username, err := ac.Svc.ResetPassword(req.Context(), email)
	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "", username)
	response.WrapRes(w, &body)
}

func (ac *AuthController) ResetPasswordVerifyHandler(w http.ResponseWriter, req *http.Request) {
	body := *response.InitRes(http.StatusOK, "", nil)
	response.WrapRes(w, &body)
}

func (ac *AuthController) ResetPasswordFinalHandler(w http.ResponseWriter, req *http.Request) {
	body := *response.InitRes(http.StatusOK, "", nil)
	response.WrapRes(w, &body)
}
