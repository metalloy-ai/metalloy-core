package user

import (
	"logiflowCore/internal/config"
	"logiflowCore/internal/database"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	GetAllUser() ([]User, []pgx.Row)
	GetUser(username string) (User, error)
}

func InitRepository(cfg config.Setting) UserRepository {
	return &Repository{db: database.GetClient(cfg)}
}