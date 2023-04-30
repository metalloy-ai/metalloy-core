package auth

import (
	"encoding/json"
	"io"

	"metalloyCore/tools"
)

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
