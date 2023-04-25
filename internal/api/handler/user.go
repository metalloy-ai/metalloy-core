package handler

import (
	"errors"
	"metalloyCore/internal/domain/user"
	"metalloyCore/pkg/response"
	"net/http"

	"github.com/uptrace/bunrouter"
)

type UserController struct {
	Repo user.UserRepository
}

func InitUserController(repo user.UserRepository) *UserController {
	return &UserController{Repo: repo}
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

	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			body := response.InitRes(http.StatusNotFound, "User not found", username)
			response.WrapRes(w, body)
		} else {
			body := response.InitRes(http.StatusInternalServerError, "Internal server error", nil)
			response.WrapRes(w, body)
		}
		return
	}

	body := response.InitRes(http.StatusOK, "", returnedUser)
	response.WrapRes(w, body)
}
