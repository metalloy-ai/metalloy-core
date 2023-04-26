package user

import "github.com/jackc/pgx/v5"

type Address struct {
    AddressID      int    `json:"address_id"`
    StreetAddress  string `json:"street_address"`
    City           string `json:"city"`
    State          string `json:"state"`
    Country        string `json:"country"`
    PostalCode     string `json:"postal_code"`
}

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
