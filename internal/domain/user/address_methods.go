package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetAddress(ctx context.Context, username string) (Address, error) {
	query := `
	SELECT a.* 
	FROM addresses as a
	JOIN users as u ON a.address_id = u.address_id
	WHERE u.username = $1`
	row := r.db.QueryRow(ctx, query, username)

	address := Address{}
	err := address.ScanFromRow(row)

	return address, err
}

func (r *Repository) CreateAddress(ctx context.Context, tx pgx.Tx, address UserCreate) (Address, error) {
	query := `
	INSERT INTO addresses (
		street_address, city, state, country, postal_code
	) VALUES (
		$1, $2, $3, $4, $5
	) returning *`
	row := tx.QueryRow(ctx, query,
		address.StreetAddress, address.City, address.State, address.Country, address.PostalCode)

	newAddress := Address{}
	err := newAddress.ScanFromRow(row)

	return newAddress, err
}

func (r *Repository) UpdateAddress(ctx context.Context, updateArr []string, args []interface{}, argsCount int, username string) (Address, error) {
	query := fmt.Sprintf(`
	UPDATE addresses
	SET %s
	WHERE address_id = (
		SELECT address_id
		FROM users
		WHERE username = $%d
	) returning *`, strings.Join(updateArr, ", "), argsCount)
	row := r.db.QueryRow(ctx, query, args...)

	newAddress := Address{}
	err := newAddress.ScanFromRow(row)

	return newAddress, err
}
