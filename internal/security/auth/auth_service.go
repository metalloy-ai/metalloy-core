package auth

import (
	"context"

	"metalloyCore/internal/domain/user"
	"metalloyCore/internal/security"
	"metalloyCore/tools"
)

type AuthService struct {
	Service user.UserService
}

func InitAuthService(service user.UserService) *AuthService {
	return &AuthService{Service: service}
}

func (as *AuthService) Login(ctx context.Context, username string, password string) (user.UserResponse, error) {
	User, err := as.Service.GetUser(ctx, username)

	if err != nil {
		return user.UserResponse{}, tools.ErrUserNotFound{}
	}

	if security.ValidatePassword(&User.Password, password) {
		return *User.ToReponse(), nil
	}

	return user.UserResponse{}, tools.ErrInvalidCredentials{}
}

func (as *AuthService) Register(ctx context.Context, newUser user.UserCreate) (user.UserResponse, error) {
	User, err := as.Service.CreateUser(ctx, newUser)

	if err != nil {
		return user.UserResponse{}, err
	}

	return User, nil
}

func (as *AuthService) ForgetPassword(ctx context.Context, username string) error {
	return nil
}