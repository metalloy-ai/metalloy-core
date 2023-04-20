package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

var ErrUserNotFound = errors.New("user not found")

func (r *Repository) GetAllUser() ([]User, []pgx.Row) {
	query := "SELECT * FROM users"
	rows, err := r.db.Query(context.Background(), query)

	if err != nil {
		return nil, nil
	}
	defer rows.Close()

	users := []User{}
	failedUsers := []pgx.Row{}

	for rows.Next() {
		user := User{}
		if err := user.ScanFromRow(rows); err != nil {

			failedUsers = append(failedUsers, rows)
		}
		users = append(users, user)
	}

	return users, failedUsers
}

func (r *Repository) GetUser(username string) (User, error) {
	query := "SELECT * FROM users WHERE username = $1"
	row := r.db.QueryRow(context.Background(), query, username)

	user := User{}

	if err := user.ScanFromRow(row); err != nil {
		if err := user.ScanFromRow(row); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return User{}, ErrUserNotFound
			}
			return User{}, err
		}
	}

	return user, nil
}