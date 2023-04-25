package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserType string

const (
	UserTypeAdmin  UserType = "admin"
	UserTypeSupplier UserType = "supplier"
	UserTypeCustomer  UserType = "customer"
)

type User struct {
	UserID           uuid.UUID  `json:"user_id"`
	Username         string     `json:"username"`
	Email            string     `json:"email"`
	Password         string     `json:"password"`
	UserType         UserType   `json:"user_type"`
	FirstName        string     `json:"first_name"`
	LastName         string     `json:"last_name"`
	PhoneNumber      string     `json:"phone_number"`
	AddressID        int        `json:"address_id"`
	RegistrationDate time.Time  `json:"registration_date"`
}

type NewUser struct {
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	UserType    UserType `json:"user_type"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	PhoneNumber string   `json:"phone_number"`
	AddressID   int      `json:"address_id"`
}

type UserResponse struct {
    UserID           uuid.UUID  `json:"user_id"`
    Username         string     `json:"username"`
    Email            string     `json:"email"`
    UserType         UserType   `json:"user_type"`
    FirstName        string     `json:"first_name"`
    LastName         string     `json:"last_name"`
    PhoneNumber      string     `json:"phone_number"`
    AddressID        int        `json:"address_id"`
    RegistrationDate time.Time  `json:"registration_date"`
}

type FullUserResponse struct {
	UserResponse
	StreetAddress  string `json:"street_address"`
	City           string `json:"city"`
	State          string `json:"state"`
	Country        string `json:"country"`
	PostalCode     string `json:"postal_code"`
}

func (u *User) ScanFromRow(row pgx.Row) error {
	err := row.Scan(
		&u.UserID,
		&u.Username,
		&u.Email,
		&u.Password,
		&u.UserType,
		&u.FirstName,
		&u.LastName,
		&u.PhoneNumber,
		&u.AddressID,
		&u.RegistrationDate,
	)
	return err
}

func (u *User) ToReponse() *UserResponse {
	return &UserResponse{
		UserID:           u.UserID,
		Username:         u.Username,
		Email:            u.Email,
		UserType:         u.UserType,
		FirstName:        u.FirstName,
		LastName:         u.LastName,
		PhoneNumber:      u.PhoneNumber,
		AddressID:        u.AddressID,
		RegistrationDate: u.RegistrationDate,
	}
}

func (ur *UserResponse) ScanFromRow(row pgx.Row) error {
	err := row.Scan(
		&ur.UserID,
		&ur.Username,
		&ur.Email,
		&ur.UserType,
		&ur.FirstName,
		&ur.LastName,
		&ur.PhoneNumber,
		&ur.AddressID,
		&ur.RegistrationDate,
	)
	return err
}

func (fr *FullUserResponse) ScanFromRow(row pgx.Row) error {
	err := row.Scan(
		&fr.UserID,
		&fr.Username,
		&fr.Email,
		&fr.UserType,
		&fr.FirstName,
		&fr.LastName,
		&fr.PhoneNumber,
		&fr.AddressID,
		&fr.RegistrationDate,
		&fr.StreetAddress,
		&fr.City,
		&fr.State,
		&fr.Country,
		&fr.PostalCode,
	)
	return err
}