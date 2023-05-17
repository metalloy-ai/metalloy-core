package auth

import (
	"context"

	"metalloyCore/internal/domain/user"
	"metalloyCore/internal/security"
	"metalloyCore/internal/security/jwt"
	"metalloyCore/tools"
)

type AuthService struct {
	Service    user.UserService
	JWTManager *jwt.JWThandler
}

func InitAuthService(service user.UserService, jwtHandler *jwt.JWThandler) *AuthService {
	return &AuthService{Service: service, JWTManager: jwtHandler}
}

func (as *AuthService) Login(ctx context.Context, username string, password string) (*AuthResponse, error) {
	User, err := as.Service.GetUser(ctx, username)

	if err != nil {
		return nil, tools.ErrUserNotFound{}
	}

	if security.ValidatePassword(&User.Password, password) {
		token, err := as.JWTManager.GenerateToken(User.UserID, User.Username, User.UserType)

		if err != nil {
			return &AuthResponse{""}, err
		}

		return &AuthResponse{token}, nil
	}

	return nil, tools.ErrInvalidCredentials{}
}

func (as *AuthService) Register(ctx context.Context, newUser *user.UserCreate) (*AuthResponse, error) {
	User, err := as.Service.CreateUser(ctx, newUser)

	if err != nil {
		return nil, err
	}

	token, err := as.JWTManager.GenerateToken(User.UserID, User.Username, User.UserType)
	if err != nil {
		return nil, err
	}
	
	return &AuthResponse{token}, nil
}

func (as *AuthService) ForgetPassword(ctx context.Context, username string) error {
	return nil
}
