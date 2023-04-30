package auth

import (
	"metalloyCore/internal/domain/user"
	"metalloyCore/internal/security"
	"metalloyCore/tools"
)

type AuthService struct {
	Repo user.UserRepository
}

func InitAuthService(repo user.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (as AuthService) Login(username string, password string) (user.UserResponse, error) {
	User, err := as.Repo.GetUser(username)

	if err != nil {
		return user.UserResponse{}, tools.ErrUserNotFound{}
	}

	if security.ValidatePassword(&User.Password, password) {
		return *User.ToReponse(), nil
	}

	return user.UserResponse{}, tools.ErrInvalidCredentials{}
}

func (as AuthService) Register(newUser user.UserCreate) (user.UserResponse, error) {
	User, err := as.Repo.CreateUser(newUser)

	if err != nil {
		return user.UserResponse{}, err
	}

	return User, nil
}
