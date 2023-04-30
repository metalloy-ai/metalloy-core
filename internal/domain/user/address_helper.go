package user

import (
	"encoding/json"
	"io"

	"github.com/jackc/pgx/v5"

	"metalloyCore/tools"
)

func (a *Address) ScanFromRow(row pgx.Row) error {
	err := row.Scan(
		&a.AddressID,
		&a.StreetAddress,
		&a.City,
		&a.State,
		&a.Country,
		&a.PostalCode,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *AddressBase) DecodeBody(data io.ReadCloser) error {
	err := json.NewDecoder(data).Decode(a)
	if err != nil {
		return tools.ErrInvalidReqBody{}
	}
	return nil
}
