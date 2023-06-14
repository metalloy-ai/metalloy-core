package twofa

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"

	"metalloyCore/tools"
)

func (lq *TwofaVerifyRequest) DecodeBody(data io.ReadCloser) error {
	err := json.NewDecoder(data).Decode(lq)
	if err != nil {
		return tools.NewBadRequestErr("Invalid JSON body: " + err.Error())
	}
	return nil
}

func (lq *TwofaVerifyRequest) Validate() error {
	if lq.Username == "" || lq.Code > 999999 || lq.Code < 100000 {
		return tools.NewBadRequestErr("Username and Code are required")
	}
	return nil
}

func (rq *TwofaResponse) DecodeBody(data io.ReadCloser) error {
	err := json.NewDecoder(data).Decode(rq)
	if err != nil {
		return tools.NewBadRequestErr("Invalid JSON body: " + err.Error())
	}
	return nil
}

func (rq *TwofaResponse) Validate() error {
	if rq.UserID == uuid.Nil || rq.Username == "" {
		return tools.NewBadRequestErr("UserID and Username are required")
	}

	return nil
}
