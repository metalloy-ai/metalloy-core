package twofa

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/google/uuid"

	"metalloyCore/internal/domain/user"
	"metalloyCore/internal/security/jwt"
	"metalloyCore/tools"
)

func SendTwofaCode(userID uuid.UUID, username string, role user.UserType, email string, serialized []byte) error {
	url := "http://172.30.16.1:2001/api/v1/auth/send"
	content := "application/json"
	body := &TwofaRequest{
		UserPayload: jwt.UserPayload{
			UserID:   userID,
			Username: username,
			Role:     role,
		},
		Email: email,
		Data:  serialized,
	}

	json, err := json.Marshal(body)
	if err != nil {
		return err
	}

	res, err := http.Post(url, content, bytes.NewBuffer(json))
	if err != nil {
		return errors.New("failed to send email: " + err.Error())
	}

	if res.StatusCode != 201 {
		return errors.New("failed to send email")
	}

	return nil
}

func VerifyTwofaCode(redisClient *redis.Client, username string, code int) (*TwofaResponse, error) {
	authPayloadJSON, err := redisClient.Get(username).Bytes()
	if err != nil {
		return nil, tools.NewUnAuthorizedErr("2fa expired")
	}

	authPayload := &TwofaResponse{}
	err = json.Unmarshal(authPayloadJSON, authPayload)
	if err != nil {
		return nil, errors.New("unable to parse payload")
	}

	if authPayload.Code != code {
		println(authPayload.Code, code)
		return nil, tools.NewUnAuthorizedErr("invalid 2fa code")
	}

	return authPayload, nil
}
