package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"metalloyCore/internal/security"
	"metalloyCore/tools"
)

type Repository struct {
	db *pgx.Conn
}

func (r *Repository) GetAllUser() ([]UserResponse, error) {
	query := `
	SELECT 
		user_id, username, email, user_type, first_name, last_name, 
		phone_number, address_id, registration_date 
	FROM users`
	rows, err := r.db.Query(context.Background(), query)

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

	if len(failedUsers) > 0 {
		return users, tools.ErrFailedUsers{FailedUsers: failedUsers}
	}

	return users, nil
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
	err := user.ScanFromRow(row)

	if errors.Is(err, pgx.ErrNoRows) {
		return FullUserResponse{}, tools.ErrUserNotFound{}
	}

	if err != nil {
		return FullUserResponse{}, err
	}

	return user, nil
}

func (r *Repository) GetUser(username string) (User, error) {
	query := `
	SELECT 
		user_id, username, email, user_type, first_name, last_name,
		phone_number, address_id, registration_date, password
	FROM users WHERE username = $1`
	row := r.db.QueryRow(context.Background(), query, username)

	user := User{}
	err := user.ScanFromRow(row)

	if errors.Is(err, pgx.ErrNoRows) {
		return User{}, tools.ErrUserNotFound{}
	}

	return user, nil
}

func (r *Repository) UpdateUser(user UserUpdate) (UserResponse, error) {
	fieldMap := map[string]interface{}{
		"email":        user.Email,
		"first_name":   user.FirstName,
		"last_name":    user.LastName,
		"phone_number": user.PhoneNumber,
	}
	updateArr, args, argsCount := tools.BuildUpdateQueryArgs(fieldMap, user.Username)

	query := fmt.Sprintf(`
	UPDATE users
	SET %s
	WHERE username = $%d
	RETURNING
		user_id, username, email, user_type, first_name, last_name,
		phone_number, address_id, registration_date`,
		strings.Join(updateArr, ", "), argsCount,
	)
	row := r.db.QueryRow(context.Background(), query, args...)

	userResponse := UserResponse{}
	err := userResponse.ScanFromRow(row)

	if errors.Is(err, pgx.ErrNoRows) {
		return UserResponse{}, tools.ErrUserNotFound{}
	}

	return userResponse, nil
}

func (r *Repository) DeleteUser(username string) error {
	query := `
	WITH deleted_users AS (
		DELETE FROM users
		WHERE username = $1
		RETURNING address_id
	  )
	DELETE FROM addresses
	WHERE address_id IN (SELECT address_id FROM deleted_users)`
	res, err := r.db.Exec(context.Background(), query, username)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return tools.ErrUserNotFound{}
		}
		println(err.Error())
		return err
	}

	if res.RowsAffected() == 0 {
		return tools.ErrUserNotFound{}
	}

	return err
}

func (r *Repository) CreateUser(user UserCreate) (UserResponse, error) {
	query := `
	INSERT INTO addresses (street_address, city, state, country, postal_code)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING *`
	row := r.db.QueryRow(context.Background(), query,
		user.StreetAddress, user.City, user.State, user.Country, user.PostalCode)

	address := Address{}
	err := address.ScanFromRow(row)

	if err != nil {
		return UserResponse{}, err
	}

	hashedPsw, err := security.HashPassword(user.Password)
	if err != nil {
		return UserResponse{}, err
	}

	query = `
	INSERT INTO users (username, email, password, user_type, first_name, last_name,
		phone_number, address_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING
		user_id, username, email, user_type, first_name, last_name,
		phone_number, address_id, registration_date`
	row = r.db.QueryRow(context.Background(), query,
		user.Username, user.Email, hashedPsw, user.UserType, user.FirstName,
		user.LastName, user.PhoneNumber, address.AddressID)

	newUser := UserResponse{}
	err = newUser.ScanFromRow(row)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return UserResponse{}, tools.ErrUserAlreadyExist{}
		}
		return UserResponse{}, err
	}

	return newUser, nil
}
