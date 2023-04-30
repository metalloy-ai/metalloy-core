package handler

import (
	"errors"
	"net/http"

	"github.com/uptrace/bunrouter"

	"metalloyCore/internal/domain/user"
	"metalloyCore/pkg/response"
	"metalloyCore/tools"
)

type UserController struct {
	Repo user.UserRepository
}

func InitUserController(repo user.UserRepository) *UserController {
	return &UserController{Repo: repo}
}

func (uc UserController) EmptyParamHandler(w http.ResponseWriter, req *http.Request) {
	body := *response.InitRes(http.StatusBadRequest, "Bad request: empty parameter", nil)
	response.WrapRes(w, &body)
}

func (uc UserController) AllUserHandler(w http.ResponseWriter, req *http.Request) {
	users, err := uc.Repo.GetAllUser()
	res := map[string]interface{}{
		"users":       users,
	}

	if err != nil {
		if errors.Is(err, tools.ErrFailedUsers{}) {
			res["failedUsers"] = err.(tools.ErrFailedUsers).Users
		}
		tools.HandleError(err, w)
		return
	}
	
	body := *response.InitRes(http.StatusOK, "", res)
	response.WrapRes(w, &body)
}

func (uc UserController) UserHandler(w http.ResponseWriter, req *http.Request) {
	username := bunrouter.ParamsFromContext(req.Context()).ByName("username")
	returnedUser, err := uc.Repo.GetFullUser(username)

	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "", returnedUser)
	response.WrapRes(w, &body)
}

func (uc UserController) UpdateUserHandler(w http.ResponseWriter, req *http.Request) {
	username := bunrouter.ParamsFromContext(req.Context()).ByName("username")
	user := *user.InitUserUpdate(username)

	err := user.DecodeBody(req.Body)
	if !tools.HandleError(err, w) {
		return
	}

	returnedUser, err := uc.Repo.UpdateUser(user)
	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "user updated success", returnedUser)
	response.WrapRes(w, &body)
}

func (uc UserController) DeleteUserHandler(w http.ResponseWriter, req *http.Request) {
	username := bunrouter.ParamsFromContext(req.Context()).ByName("username")
	
	err := uc.Repo.DeleteUser(username)
	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "user deleted success", nil)
	response.WrapRes(w, &body)
}

func (uc UserController) GetAddressHandler(w http.ResponseWriter, req *http.Request) {
	username := bunrouter.ParamsFromContext(req.Context()).ByName("username")
	returnedAddress, err := uc.Repo.GetAddress(username)

	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "", returnedAddress)
	response.WrapRes(w, &body)
}

func (uc UserController) UpdateAddressHandler(w http.ResponseWriter, req *http.Request) {
	username := bunrouter.ParamsFromContext(req.Context()).ByName("username")
	address := user.AddressBase{}
	
	err := address.DecodeBody(req.Body)
	if !tools.HandleError(err, w) {
		return
	}

	returnedAddress, err := uc.Repo.UpdateAddress(address, username)
	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "address updated success", returnedAddress)
	response.WrapRes(w, &body)
}
