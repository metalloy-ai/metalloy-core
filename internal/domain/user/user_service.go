package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"

	"metalloyCore/internal/security"
	"metalloyCore/tools"
)

type UserService interface {
	GetAllUser(ctx context.Context) ([]UserResponse, error)
	GetFullUser(ctx context.Context, username string) (FullUserResponse, error)
	GetUser(ctx context.Context, username string) (User, error)
	UpdateUser(ctx context.Context, user UserUpdate) (UserResponse, error)
	CreateUser(ctx context.Context, newUser UserCreate) (UserResponse, error)
	DeleteUser(ctx context.Context, username string) error
	GetAddress(ctx context.Context, username string) (Address, error)
	UpdateAddress(ctx context.Context, address AddressBase, username string) (Address, error)
}

type Service struct {
	Repo UserRepository
}

func InitUserService(repo UserRepository) UserService {
	return &Service{Repo: repo}
}

func (us *Service) GetAllUser(ctx context.Context) ([]UserResponse, error) {
	users, failedUsers := us.Repo.GetAllUser(ctx)

	if len(failedUsers) > 0 {
		return nil, tools.ErrFailedUsers{FailedUsers: failedUsers}
	}

	return users, nil
}

func (us *Service) GetFullUser(ctx context.Context, username string) (FullUserResponse, error) {
	user, err := us.Repo.GetFullUser(ctx, username)

	handledUser, err := tools.HandleEmptyError(user, err)
	return handledUser.(FullUserResponse), err
}

func (us *Service) GetUser(ctx context.Context, username string) (User, error) {
	user, err := us.Repo.GetUser(ctx, username)

	handledUser, err := tools.HandleEmptyError(user, err)
	return handledUser.(User), err
}

func (us *Service) UpdateUser(ctx context.Context, user UserUpdate) (UserResponse, error) {
	fieldMap := map[string]interface{}{
		"email":        user.Email,
		"first_name":   user.FirstName,
		"last_name":    user.LastName,
		"phone_number": user.PhoneNumber,
	}
	updateArr, args, argsCount := tools.BuildUpdateQueryArgs(fieldMap, user.Username)
	updatedUser, err := us.Repo.UpdateUser(ctx, updateArr, args, argsCount)

	handledUpdatedUser, err := tools.HandleEmptyError(updatedUser, err)
	return handledUpdatedUser.(UserResponse), err
}

func (us *Service) CreateUser(ctx context.Context, newUser UserCreate) (UserResponse, error) {
	hashedPsw, err := security.HashPassword(newUser.Password)
	if err != nil {
		return UserResponse{}, err
	}

	user, err := us.Repo.CreateUser(ctx, newUser, hashedPsw)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return UserResponse{}, tools.ErrUserAlreadyExist{}
		}

		return UserResponse{}, err
	}

	return user, nil
}

func (us *Service) DeleteUser(ctx context.Context, username string) error {
	if err := us.Repo.DeleteUser(ctx, username); err != nil {
		return err
	}

	return nil
}
