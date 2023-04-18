package user

import (
	"github.com/google/uuid"
	"time"
)

type UserType string

const (
	UserTypeManufacturer UserType = "manufacturer"
	UserTypedistributor  UserType = "distributor"
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
	Address          string     `json:"address"`
	RegistrationDate time.Time  `json:"registration_date"`
	RoleID           int64      `json:"role_id"`
}

type Role struct {
	RoleID      int64  `json:"role_id"`
	RoleName    string `json:"role_name"`
	Description string `json:"description"`
}