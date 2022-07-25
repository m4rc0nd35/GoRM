package users

import (
	"context"

	"github.com/m4rc0nd35/db-golang/entities"
)

type IUserRepository interface {
	Create(ctx context.Context, newUser entities.User) error
	GetAll(ctx context.Context) ([]entities.User, error)
	GetByID(ctx context.Context, ID int64) (*entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
}
