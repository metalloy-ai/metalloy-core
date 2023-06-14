package auth

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/go-redis/redis"

	"metalloyCore/internal/domain/auth/twofa"
	"metalloyCore/internal/domain/user"
	"metalloyCore/internal/security"
	"metalloyCore/internal/security/jwt"
	"metalloyCore/tools"
)

type AuthService interface {
	Login(ctx context.Context, username string, password string) error
	LoginVerify(ctx context.Context, username string, code int) (*AuthResponse, error)
	Register(ctx context.Context, newUser *user.UserCreate) error
	RegisterVerify(ctx context.Context, username string, code int) (*AuthResponse, error)
	ResetPassword(ctx context.Context, email string) (string, error)
	ResetPasswordVerify(ctx context.Context, username string, code int) (*twofa.TwofaResponse, error)
	ResetPasswordFinal(ctx context.Context, username string, password string) (*user.UserResponse, error)
}

type Service struct {
	Service     user.UserService
	JWTManager  *jwt.JWThandler
	RedisClient *redis.Client
}

func InitAuthService(service user.UserService, jwtHandler *jwt.JWThandler, redisClient *redis.Client) AuthService {
	return &Service{Service: service, JWTManager: jwtHandler, RedisClient: redisClient}
}

func (as *Service) Login(ctx context.Context, username string, password string) error {
	User, err := as.Service.GetUser(ctx, username)
	if err != nil {
		return tools.ErrUserNotFound{}
	}

	if security.ValidatePassword(&User.Password, password) {
		return twofa.SendTwofaCode(User.UserID, User.Username, User.UserType, User.Email, nil)
	}

	return tools.NewUnAuthorizedErr("invalid password")
}

func (as *Service) LoginVerify(ctx context.Context, username string, code int) (*AuthResponse, error) {
	authPayload, err := twofa.VerifyTwofaCode(as.RedisClient, username, code)
	if err != nil {
		return nil, err
	}

	jwtToken, err := as.JWTManager.GenerateToken(authPayload.UserID, authPayload.Username, authPayload.Role)
	if err != nil {
		return &AuthResponse{""}, err
	}

	return &AuthResponse{jwtToken}, nil
}

func (as *Service) Register(ctx context.Context, newUser *user.UserCreate) error {
	serialized, err := json.Marshal(newUser)
	if err != nil {
		return err
	}

	return twofa.SendTwofaCode(newUser.UserID, newUser.Username, newUser.UserType, newUser.Email, serialized)
}

func (as *Service) RegisterVerify(ctx context.Context, username string, code int) (*AuthResponse, error) {
	twofaResponse, err := twofa.VerifyTwofaCode(as.RedisClient, username, code)
	if err != nil {
		return nil, err
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(twofaResponse.Data)
	if err != nil {
		return nil, err
	}

	newUser := &user.UserCreate{}
	err = json.Unmarshal(decodedBytes, &newUser)
	if err != nil {
		return nil, err
	}

	hashed, err := security.HashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}
	newUser.Password = hashed
	
	user, err := as.Service.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	jwtToken, err := as.JWTManager.GenerateToken(user.UserID, user.Username, user.UserType)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{jwtToken}, nil
}

func (as *Service) ResetPassword(ctx context.Context, email string) (string, error) {
	user, err := as.Service.GetUserByEmail(ctx, email)
	if err != nil {
		println(user)
		return "", tools.ErrUserNotFound{}
	}

	err = twofa.SendTwofaCode(user.UserID, user.Username, user.UserType, user.Email, nil)
	if err != nil {
		return "", err
	}

	return user.Username, nil
}

func (as *Service) ResetPasswordVerify(ctx context.Context, username string, code int) (*twofa.TwofaResponse, error) {
	twofaResponse, err := twofa.VerifyTwofaCode(as.RedisClient, username, code)
	if err != nil {
		return nil, err
	}

	return twofaResponse, nil
}

func (as *Service) ResetPasswordFinal(ctx context.Context, username string, password string) (*user.UserResponse, error) {
	hashed, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}

	User, err := as.Service.UpdateUserPassword(ctx, username, hashed)
	if err != nil {
		return nil, err
	}

	return User, nil
}
