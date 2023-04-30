package user

import (
	"time"

	"github.com/google/uuid"
)

type UserType string

const (
	UserTypeAdmin    UserType = "admin"
	UserTypeSupplier UserType = "supplier"
	UserTypeCustomer UserType = "customer"
)

type UserBase struct {
	UserID           uuid.UUID `json:"user_id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	UserType         UserType  `json:"user_type"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	PhoneNumber      string    `json:"phone_number"`
	AddressID        int       `json:"address_id"`
	RegistrationDate time.Time `json:"registration_date"`
}

type User struct {
	UserBase
	Password string `json:"password"`
}

type UserResponse struct {
	UserBase
}

type FullUserResponse struct {
	UserBase
	AddressBase
}

type UserUpdate struct {
	UserBase
}

type UserCreate struct {
	FullUserResponse
	Password string `json:"password"`
}

type AddressBase struct {
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	PostalCode    string `json:"postal_code"`
}

type Address struct {
	AddressID int `json:"address_id"`
	AddressBase
}