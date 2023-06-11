package auth

import (
	"encoding/json"
	"io"

	"metalloyCore/tools"
)

func (lq *LoginRequest) DecodeBody(data io.ReadCloser) error {
	err := json.NewDecoder(data).Decode(lq)
	if err != nil {
		return tools.NewBadRequestErr("Invalid JSON body: " + err.Error())
	}
	return nil
}

func (lq *LoginRequest) Validate() error {
	if lq.Username == "" || lq.Password == "" {
		return tools.NewBadRequestErr("Username and password are required")
	}
	return nil
}
