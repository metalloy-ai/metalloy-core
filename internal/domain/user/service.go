package user

import (
	"context"
	"errors"
	"metalloyCore/tools"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

func (r *Repository) GetAllUser() ([]UserResponse, []pgx.Row) {
	query := `
	SELECT 
		user_id, username, email, user_type, first_name, last_name, 
		phone_number, address_id, registration_date 
	FROM users`
	rows, err := r.db.Query(context.Background(), query)

	if err != nil {
		return nil, nil
	}
	defer rows.Close()

	users := []UserResponse{}
	failedUsers := []pgx.Row{}

	for rows.Next() {
		user := UserResponse{}
		
		if err := user.ScanFromRow(rows); err != nil {
			failedUsers = append(failedUsers, rows)
		} else {
			users = append(users, user)
		}

	}

	return users, failedUsers
}

func (r *Repository) GetFullUser(username string) (FullUserResponse, error) {
	query := `
	SELECT 
		u.user_id, u.username, u.email, u.user_type, u.first_name, u.last_name, 
		u.phone_number, u.address_id, u.registration_date,
		a.street_address, a.city, a.state, a.country, a.postal_code
	FROM users as u
	JOIN addresses as a ON u.address_id = a.address_id
	WHERE u.username = $1`
	row := r.db.QueryRow(context.Background(), query, username)

	user := FullUserResponse{}
	err := user.ScanFromRow(row);

	if errors.Is(err, pgx.ErrNoRows) {
		return FullUserResponse{}, tools.ErrUserNotFound
	}

	if err != nil {
		return FullUserResponse{}, err
	}

	return user, nil
}

func (r *Repository) GetUser(username string) (User, error) {
	query := "SELECT * FROM users WHERE username = $1"
	row := r.db.QueryRow(context.Background(), query, username)

	user := User{}
	err := user.ScanFromRow(row);

	if errors.Is(err, pgx.ErrNoRows) {
		return User{}, tools.ErrUserNotFound
	}

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *Repository) UpdateUser(user User) (UserResponse, error) {return UserResponse{}, nil}

func (r *Repository) DeleteUser(username string) error {return nil}

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
		return Address{}, tools.ErrUserNotFound
	}

	return address, nil
}

func (r *Repository) UpdateAddress(Address Address) error {return nil}