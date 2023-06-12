package user

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"metalloyCore/internal/config"
	"metalloyCore/internal/database"
)

type UserRepository interface {
	GetAllUser(ctx context.Context, pageIdx string, pageSize int) ([]*UserResponse, []pgx.Row)
	GetFullUser(ctx context.Context, username string) (*FullUserResponse, error)
	GetUser(ctx context.Context, username string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, updateArr []string, args []interface{}, argsCount int) (*UserResponse, error)
	UpdateUserPassword(ctx context.Context, username string, password string) (*UserResponse, error)
	CreateUser(ctx context.Context, newUser *UserCreate, hashedPsw string) (*UserResponse, error)
	DeleteUser(ctx context.Context, username string) error
	GetAddress(ctx context.Context, username string) (*Address, error)
	UpdateAddress(ctx context.Context, updateArr []string, args []interface{}, argsCount int, username string) (*Address, error)
}

type Repository struct {
	db *pgxpool.Pool
}

func InitRepository(cfg config.Setting) UserRepository {
	return &Repository{db: database.GetPool(cfg)}
}
