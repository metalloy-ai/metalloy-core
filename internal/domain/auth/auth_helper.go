package auth

import (
	"encoding/json"
	"io"

	"metalloyCore/tools"
)

func (lq *LoginRequest) DecodeBody(data io.ReadCloser) error {
	err := json.NewDecoder(data).Decode(lq)
	if err != nil {
		return tools.ErrInvalidReq{}
	}
	return nil
}

func (lq *LoginRequest) Validate() error {
	if lq.Username == "" || lq.Password == "" {
		return tools.ErrMissingParams{}
	}
	return nil
}
