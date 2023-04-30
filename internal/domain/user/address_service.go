package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"metalloyCore/tools"
)

func (r *Repository) GetAddress(username string) (Address, error) {
	query := `
	SELECT a.* 
	FROM addresses as a
	JOIN users as u ON a.address_id = u.address_id
	WHERE u.username = $1`
	row := r.db.QueryRow(context.Background(), query, username)

	address := Address{}
	err := address.ScanFromRow(row)

	if errors.Is(err, pgx.ErrNoRows) {
		return Address{}, tools.ErrUserNotFound{}
	}

	return address, nil
}

func (r *Repository) CreateAddress(tx pgx.Tx, address UserCreate) (Address, error) {
	query := `
	INSERT INTO addresses (
		street_address, city, state, country, postal_code
	) VALUES (
		$1, $2, $3, $4, $5
	) returning *`
	row := tx.QueryRow(context.Background(), query,
		address.StreetAddress, address.City, address.State, address.Country, address.PostalCode)

	newAddress := Address{}
	err := newAddress.ScanFromRow(row)
	return newAddress, err
}

func (r *Repository) UpdateAddress(address AddressBase, username string) (Address, error) {
	fieldMap := map[string]interface{}{
		"street_address": address.StreetAddress,
		"city":           address.City,
		"state":          address.State,
		"country":        address.Country,
		"postal_code":    address.PostalCode,
	}
	updateArr, args, argsCount := tools.BuildUpdateQueryArgs(fieldMap, username)

	query := fmt.Sprintf(`
	UPDATE addresses
	SET %s
	WHERE address_id = (
		SELECT address_id
		FROM users
		WHERE username = $%d
	) returning *`, strings.Join(updateArr, ", "), argsCount)
	row := r.db.QueryRow(context.Background(), query, args...)

	newAddress := Address{}
	err := newAddress.ScanFromRow(row)

	if errors.Is(err, pgx.ErrNoRows) {
		return Address{}, tools.ErrUserNotFound{}
	}

	return newAddress, nil
}
