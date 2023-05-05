package user

import (
	"errors"

	"github.com/jackc/pgx/v5"

	"metalloyCore/tools"
)

type UserService interface {
	GetAllUser() ([]UserResponse, error)
	GetFullUser(username string) (FullUserResponse, error)
	GetUser(username string) (User, error)
	UpdateUser(user UserUpdate) (UserResponse, error)
	CreateUser(newUser UserCreate) (UserResponse, error)
	DeleteUser(username string) error
	GetAddress(username string) (Address, error)
	UpdateAddress(address AddressBase, username string) (Address, error)
}

type Service struct {
	Repo UserRepository
}

func InitUserService(repo UserRepository) UserService {
	return &Service{Repo: repo}
}

func (us *Service) GetAllUser() ([]UserResponse, error) {
	users, failedUsers := us.Repo.GetAllUser()

	if len(failedUsers) > 0 {
		return users, tools.ErrFailedUsers{FailedUsers: failedUsers}
	}

	return users, nil
}

func (us *Service) GetFullUser(username string) (FullUserResponse, error) {
	user, err := us.Repo.GetFullUser(username)

	if errors.Is(err, pgx.ErrNoRows) {
		return FullUserResponse{}, tools.ErrUserNotFound{}
	}
	if err != nil {
		return FullUserResponse{}, err
	}

	return user, nil
}

func (us *Service) GetUser(username string) (User, error) {
	return us.Repo.GetUser(username)
}

func (us *Service) UpdateUser(user UserUpdate) (UserResponse, error) {
	return us.Repo.UpdateUser(user)
}

func (us *Service) CreateUser(newUser UserCreate) (UserResponse, error) {
	return us.Repo.CreateUser(newUser)
}

func (us *Service) DeleteUser(username string) error {
	return us.Repo.DeleteUser(username)
}
