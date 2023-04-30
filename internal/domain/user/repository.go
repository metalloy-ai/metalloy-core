package user

import (
	"metalloyCore/internal/config"
	"metalloyCore/internal/database"
)

type UserRepository interface {
	GetAllUser() ([]UserResponse, error)
	GetFullUser(username string) (FullUserResponse, error)
	GetUser(username string) (User, error)
	UpdateUser(user UserUpdate) (UserResponse, error)
	CreateUser(newUser UserCreate) (UserResponse, error)
	DeleteUser(username string) error
	GetAddress(username string) (Address, error)
	UpdateAddress(address AddressBase, username string) (Address, error)
}

func InitRepository(cfg config.Setting) UserRepository {
	return &Repository{db: database.GetClient(cfg)}
}
