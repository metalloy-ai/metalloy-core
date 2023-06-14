package user

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	GetAllUser(ctx context.Context, pageIdx string, pageSize int) ([]*UserResponse, []pgx.Row)
	GetFullUser(ctx context.Context, username string) (*FullUserResponse, error)
	GetUser(ctx context.Context, username string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, updateArr []string, args []interface{}, argsCount int) (*UserResponse, error)
	UpdateUserPassword(ctx context.Context, username string, password string) (*UserResponse, error)
	CreateUser(ctx context.Context, newUser *UserCreate) (*UserResponse, error)
	DeleteUser(ctx context.Context, username string) error
	GetAddress(ctx context.Context, username string) (*Address, error)
	UpdateAddress(ctx context.Context, updateArr []string, args []interface{}, argsCount int, username string) (*Address, error)
}

type UserService interface {
	GetAllUser(ctx context.Context, pageIdx string, sizeRaw string) ([]*UserResponse, error)
	GetFullUser(ctx context.Context, username string) (*FullUserResponse, error)
	GetUser(ctx context.Context, username string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, user *UserUpdate) (*UserResponse, error)
	UpdateUserPassword(ctx context.Context, username string, password string) (*UserResponse, error)
	CreateUser(ctx context.Context, newUser *UserCreate) (*UserResponse, error)
	DeleteUser(ctx context.Context, username string) error
	GetAddress(ctx context.Context, username string) (*Address, error)
	UpdateAddress(ctx context.Context, address *AddressBase, username string) (*Address, error)
}