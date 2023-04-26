package auth

import (
	"encoding/json"
	"io"

	"metalloyCore/tools"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (lq *LoginReq) DecodeBody(data io.ReadCloser) error {
	err := json.NewDecoder(data).Decode(lq)
	if err != nil {
		return tools.ErrInvalidReqBody{}
	}
	return nil
}

func (lq *LoginReq) Validate() error {
	if lq.Username == "" || lq.Password == "" {
		return tools.ErrMissingParams{}
	}
	return nil
}