package handler

import (
	"errors"
	"net/http"

	"github.com/uptrace/bunrouter"

	"metalloyCore/internal/domain/user"
	"metalloyCore/pkg/response"
	"metalloyCore/pkg/validator"
	"metalloyCore/tools"
)

type UserController struct {
	Service user.UserService
}

func InitUserController(service user.UserService) *UserController {
	return &UserController{Service: service}
}

func (uc *UserController) AllUserHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	username, sizeRaw := req.URL.Query().Get("page-idx"), req.URL.Query().Get("page-size")
	users, err := uc.Service.GetAllUser(ctx, username, sizeRaw)
	res := map[string]interface{}{"users": users}

	if err != nil {
		if errors.Is(err, tools.ErrFailedUsers{}) {
			res["failedUsers"] = err.(tools.ErrFailedUsers).Users
		}
		tools.HandleError(err, w)
		return
	}

	body := *response.InitRes(http.StatusOK, "", &res)
	response.WrapRes(w, &body)
}

func (uc *UserController) UserHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	username := bunrouter.ParamsFromContext(ctx).ByName("username")

	err := validator.ValidatePayload(req, username)
	if !tools.HandleError(err, w) {
		return
	}

	returnedUser, err := uc.Service.GetFullUser(ctx, username)
	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "", *returnedUser)
	response.WrapRes(w, &body)
}

func (uc *UserController) UpdateUserHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	username := bunrouter.ParamsFromContext(ctx).ByName("username")

	err := validator.ValidatePayload(req, username)
	if !tools.HandleError(err, w) {
		return
	}

	user := *user.InitUserUpdate(username)
	err = user.DecodeBody(req.Body)
	if !tools.HandleError(err, w) {
		return
	}

	returnedUser, err := uc.Service.UpdateUser(ctx, &user)
	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "user updated success", *returnedUser)
	response.WrapRes(w, &body)
}

func (uc *UserController) DeleteUserHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	username := bunrouter.ParamsFromContext(ctx).ByName("username")

	err := uc.Service.DeleteUser(ctx, username)
	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "user deleted success", nil)
	response.WrapRes(w, &body)
}

func (uc *UserController) GetAddressHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	username := bunrouter.ParamsFromContext(ctx).ByName("username")

	err := validator.ValidatePayload(req, username)
	if !tools.HandleError(err, w) {
		return
	}

	returnedAddress, err := uc.Service.GetAddress(ctx, username)
	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "", *returnedAddress)
	response.WrapRes(w, &body)
}

func (uc *UserController) UpdateAddressHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	username := bunrouter.ParamsFromContext(ctx).ByName("username")

	err := validator.ValidatePayload(req, username)
	if !tools.HandleError(err, w) {
		return
	}

	address := &user.AddressBase{}
	err = address.DecodeBody(req.Body)
	if !tools.HandleError(err, w) {
		return
	}

	returnedAddress, err := uc.Service.UpdateAddress(ctx, address, username)
	if !tools.HandleError(err, w) {
		return
	}

	body := *response.InitRes(http.StatusOK, "address updated success", *returnedAddress)
	response.WrapRes(w, &body)
}
