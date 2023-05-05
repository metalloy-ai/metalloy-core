package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"

	"metalloyCore/tools"
)

func (r *Repository) GetAllUser(ctx context.Context) ([]UserResponse, []pgx.Row) {
	query := `
	SELECT 
		user_id, username, email, user_type, first_name, last_name, 
		phone_number, address_id, registration_date 
	FROM users`
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return []UserResponse{}, nil
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

func (r *Repository) GetFullUser(ctx context.Context, username string) (FullUserResponse, error) {
	query := `
	SELECT 
		u.user_id, u.username, u.email, u.user_type, u.first_name, u.last_name, 
		u.phone_number, u.address_id, u.registration_date,
		a.street_address, a.city, a.state, a.country, a.postal_code
	FROM users as u
	JOIN addresses as a ON u.address_id = a.address_id
	WHERE u.username = $1`
	row := r.db.QueryRow(ctx, query, username)

	user := FullUserResponse{}
	err := user.ScanFromRow(row)

	return user, err
}

func (r *Repository) GetUser(ctx context.Context, username string) (User, error) {
	query := `
	SELECT 
		user_id, username, email, user_type, first_name, last_name,
		phone_number, address_id, registration_date, password
	FROM users WHERE username = $1`
	row := r.db.QueryRow(ctx, query, username)

	user := User{}
	err := user.ScanFromRow(row)

	return user, err
}

func (r *Repository) UpdateUser(ctx context.Context, updateArr []string, args []interface{}, argsCount int) (UserResponse, error) {
	query := fmt.Sprintf(`
	UPDATE users
	SET %s
	WHERE username = $%d
	RETURNING
		user_id, username, email, user_type, first_name, last_name,
		phone_number, address_id, registration_date`,
		strings.Join(updateArr, ", "), argsCount,
	)
	row := r.db.QueryRow(ctx, query, args...)

	userResponse := UserResponse{}
	err := userResponse.ScanFromRow(row)

	return userResponse, err
}

func (r *Repository) CreateUser(ctx context.Context, user UserCreate, hashedPsw string) (UserResponse, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return UserResponse{}, err
	}
	defer tx.Rollback(ctx)

	address, err := r.CreateAddress(ctx, tx, user)

	if err != nil {
		return UserResponse{}, err
	}

	query := `
        INSERT INTO users (username, email, password, user_type, first_name, last_name,
            phone_number, address_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING
            user_id, username, email, user_type, first_name, last_name,
            phone_number, address_id, registration_date`
	row := tx.QueryRow(ctx, query,
		user.Username, user.Email, hashedPsw, user.UserType, user.FirstName,
		user.LastName, user.PhoneNumber, address.AddressID)

	newUser := UserResponse{}
	err = newUser.ScanFromRow(row)

	if err != nil {
		return UserResponse{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return UserResponse{}, err
	}

	return newUser, err
}

func (r *Repository) DeleteUser(ctx context.Context, username string) error {
	query := `
	WITH deleted_users AS (
		DELETE FROM users
		WHERE username = $1
		RETURNING address_id
	)
	DELETE FROM addresses
	WHERE address_id IN (SELECT address_id FROM deleted_users)`
	res, err := r.db.Exec(ctx, query, username)

	if res.RowsAffected() == 0 {
		return tools.ErrUserNotFound{}
	}

	if err != nil {
		return err
	}

	return err
}
