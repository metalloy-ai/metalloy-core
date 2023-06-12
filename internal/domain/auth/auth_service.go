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
		return twofa.SendTwofaCode(User.UserID, User.Username, User.UserType, User.Email, nil)
	}

	return tools.NewUnAuthorizedErr("invalid password")
}

func (as *AuthService) LoginVerify(ctx context.Context, username string, code int) (*AuthResponse, error) {
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

func (as *AuthService) Register(ctx context.Context, newUser *user.UserCreate) error {
	serialized, err := json.Marshal(newUser)
	if err != nil {
		return err
	}

	return twofa.SendTwofaCode(newUser.UserID, newUser.Username, newUser.UserType, newUser.Email, serialized)
}

func (as *AuthService) RegisterVerify(ctx context.Context, username string, code int) (*AuthResponse, error) {
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

func (as *AuthService) ResetPassword(ctx context.Context, email string) (string, error) {
	User, err := as.Service.GetUserByEmail(ctx, email)
	if err != nil {
		println(User)
		return "", tools.ErrUserNotFound{}
	}

	err = twofa.SendTwofaCode(User.UserID, User.Username, User.UserType, User.Email, nil)
	if err != nil {
		return "", err
	}

	return User.Username, nil
}

func (as *AuthService) ResetPasswordVerify(ctx context.Context, username string, code int) error {
	return nil
}

func (as *AuthService) ResetPasswordFinal(ctx context.Context, username string, password string) error {
	return nil
}