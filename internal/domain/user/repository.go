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
}

func InitRepository(cfg config.Setting) UserRepository {
	return &Repository{db: database.GetClient(cfg)}
}
