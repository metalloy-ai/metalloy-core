package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	"metalloyCore/internal/config"
	"metalloyCore/internal/domain/user"
	"metalloyCore/tools"
)

type JWThandler struct {
	SecretKey string
	TokenLife time.Duration
}

type Claims struct {
	jwt.StandardClaims
	UserPayload
}

type UserPayload struct {
	UserID   uuid.UUID     `json:"user_id"`
	Username string        `json:"username"`
	Role     user.UserType `json:"role"`
}

func InitJWTHandler(cfg config.Setting) *JWThandler {
	return &JWThandler{cfg.JwtKey, time.Hour * time.Duration(cfg.TokenDuration)}
}

func (j *JWThandler) GenerateToken(userID uuid.UUID, username string, role user.UserType) (string, error) {
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.TokenLife).Unix(),
		},
		UserPayload: UserPayload{
			UserID:   userID,
			Username: username,
			Role:     role,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	return token.SignedString([]byte(j.SecretKey))
}

func (j *JWThandler) RefreshToken(tokenInput string) (string, error) {
	claims, err := j.ValidateToken(tokenInput)
	if err != nil {
		return "", err
	}

	return j.GenerateToken(claims.UserID, claims.Username, claims.Role)
}

func (j *JWThandler) ValidateToken(tokenInput string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenInput,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		if err, ok := err.(*jwt.ValidationError); ok && err.Errors == jwt.ValidationErrorExpired {
			return nil, tools.ErrExpiredToken{}
		}
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, tools.ErrParseClaims{}
	}

	return claims, nil
}

func (j *JWThandler) ValidateRequest(r *http.Request) (*Claims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, tools.ErrNoAuthHeader{}
	}

	header := strings.Split(authHeader, " ")
	if len(header) != 2 || strings.ToLower(header[0]) != "bearer" {
		return nil, tools.ErrInvalidAuthHeader{}
	}

	return j.ValidateToken(header[1])
}
