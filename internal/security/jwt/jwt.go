package jwt

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	"metalloyCore/internal/config"
	"metalloyCore/internal/domain/user"
)

type JWThandler struct {
	SecretKey string
	TokenLife time.Duration
}

type Claims struct {
	jwt.StandardClaims
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
		UserID:   userID,
		Username: username,
		Role:     role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	return token.SignedString([]byte(j.SecretKey))
}

func (j *JWThandler) ValidateToken(tokenInput string) (*Claims, error) {
	return nil, nil
}

func (j *JWThandler) ValidateRequest(r *http.Request) (*Claims, error) {
	return nil, nil
}
