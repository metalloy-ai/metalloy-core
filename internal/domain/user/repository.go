package user

import (
	"github.com/jackc/pgx/v5"

	"metalloyCore/internal/config"
	"metalloyCore/internal/database"
)

type UserRepository interface {
	GetAllUser() ([]UserResponse, []pgx.Row)
	GetFullUser(username string) (FullUserResponse, error)
	GetUser(username string) (User, error)
	UpdateUser(user UserUpdate) (UserResponse, error)
	CreateUser(newUser UserCreate) (UserResponse, error)
	DeleteUser(username string) error
	GetAddress(username string) (Address, error)
	UpdateAddress(address AddressBase, username string) (Address, error)
}

type Repository struct {
	db *pgx.Conn
}

func InitRepository(cfg config.Setting) UserRepository {
	return &Repository{db: database.GetClient(cfg)}
}
