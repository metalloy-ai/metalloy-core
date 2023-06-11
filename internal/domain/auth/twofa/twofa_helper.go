package twofa

import (
	"encoding/json"
	"io"

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
