package auth

import (
	"metalloyCore/internal/domain/user"
	"metalloyCore/internal/security"
)

type AuthService struct {
	Repo user.UserRepository
}

func InitAuthService(repo user.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (as AuthService) Login(username string, password string) (user.User, error) {
	User, err := as.Repo.GetUser(username)

	if err != nil {
		return user.User{}, user.ErrUserNotFound
	}

	if security.ValidatePassword(&User.Password, password) {
		return User, nil
	}

	return user.User{}, ErrInvalidCredentials
}

func (as AuthService) Register(username string, password string) {}
