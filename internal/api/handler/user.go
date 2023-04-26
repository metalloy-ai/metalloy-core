package handler

import (
	"net/http"

	"metalloyCore/internal/domain/user"
	"metalloyCore/pkg/response"
	"metalloyCore/tools"

	"github.com/uptrace/bunrouter"
)

type UserController struct {
	Repo user.UserRepository
}

func InitUserController(repo user.UserRepository) *UserController {
	return &UserController{Repo: repo}
}

func (uc UserController) EmptyParamHandler(w http.ResponseWriter, req *http.Request) {
	body := response.InitRes(http.StatusBadRequest, "Bad request: empty parameter", nil)
	response.WrapRes(w, body)
}

func (uc UserController) AllUserHandler(w http.ResponseWriter, req *http.Request) {
	users, failedUser := uc.Repo.GetAllUser()
	res := map[string]interface{}{
		"users":       users,
		"failed_user": failedUser,
	}

	body := response.InitRes(http.StatusOK, "", res)
	response.WrapRes(w, body)
}

func (uc UserController) UserHandler(w http.ResponseWriter, req *http.Request) {
	params := bunrouter.ParamsFromContext(req.Context())
	username := params.ByName("username")
	returnedUser, err := uc.Repo.GetFullUser(username)

	if !tools.HandleError(err, w) {
		return
	}

	body := response.InitRes(http.StatusOK, "", returnedUser)
	response.WrapRes(w, body)
}

func (uc UserController) UpdateUserHandler(w http.ResponseWriter, req *http.Request) {}

func (uc UserController) DeleteUserHandler(w http.ResponseWriter, req *http.Request) {}

func (uc UserController) GetAddressHandler(w http.ResponseWriter, req *http.Request) {
	params := bunrouter.ParamsFromContext(req.Context())
	username := params.ByName("username")
	returnedAddress, err := uc.Repo.GetAddress(username)

	if !tools.HandleError(err, w) {
		return
	}

	body := response.InitRes(http.StatusOK, "", returnedAddress)
	response.WrapRes(w, body)
}

func (uc UserController) UpdateAddressHandler(w http.ResponseWriter, req *http.Request) {}