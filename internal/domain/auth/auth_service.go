package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-redis/redis"

	"metalloyCore/internal/domain/user"
	"metalloyCore/internal/security"
	"metalloyCore/internal/security/jwt"
	"metalloyCore/tools"
)

type AuthService struct {
	Service     user.UserService
	JWTManager  *jwt.JWThandler
	RedisClient *redis.Client
}

func InitAuthService(service user.UserService, jwtHandler *jwt.JWThandler, redisClient *redis.Client) *AuthService {
	return &AuthService{Service: service, JWTManager: jwtHandler, RedisClient: redisClient}
}

func (as *AuthService) Login(ctx context.Context, username string, password string) error {
	User, err := as.Service.GetUser(ctx, username)
	if err != nil {
		return tools.ErrUserNotFound{}
	}

	if security.ValidatePassword(&User.Password, password) {
		url := "http://172.30.16.1:2001/api/v1/auth/send"
		content := "application/json"
		body := &AuthGenRequest{
			AuthPayload: AuthPayload{
				UserID:   User.UserID,
				Username: User.Username,
				UserType: User.UserType,
			},
			Email: User.Email,
		}

		json, err := json.Marshal(body)
		if err != nil {
			return err
		}

		res, err := http.Post(url, content, bytes.NewBuffer(json))
		if err != nil {
			return err
		}

		if res.StatusCode != 201 {
			return errors.New("failed to send email")
		}

		return nil
	}

	return tools.ErrInvalidCredentials{}
}

func (as *AuthService) LoginVerify(ctx context.Context, token string) (*AuthResponse, error) {
	authPayloadJSON, err := as.RedisClient.Get(token).Bytes()
	if err != nil {
		return nil, errors.New("invalid token")
	}

	authPayload := &AuthPayload{}
	err = json.Unmarshal(authPayloadJSON, authPayload)

	if err != nil {
		return nil, err
	}

	jwtToken, err := as.JWTManager.GenerateToken(authPayload.UserID, authPayload.Username, authPayload.UserType)
	if err != nil {
		return &AuthResponse{""}, err
	}

	return &AuthResponse{jwtToken}, nil
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
