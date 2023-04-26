package user

import (
	"metalloyCore/internal/config"
	"metalloyCore/internal/database"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	GetAllUser() ([]UserResponse, []pgx.Row)
	GetFullUser(username string) (FullUserResponse, error)
	GetUser(username string) (User, error)
	UpdateUser(user User) (UserResponse, error)
	DeleteUser(username string) error
	GetAddress(username string) (Address, error)
	UpdateAddress(address Address) error
}

func InitRepository(cfg config.Setting) UserRepository {
	return &Repository{db: database.GetClient(cfg)}
}
