package repository

import (
	"context"

	"github.com/mar-cial/pizza-auth/internal/domain"
)

type AuthUsers interface {
	CreateUser(ctx context.Context, user domain.User) error
	UpdateUser(ctx context.Context, user domain.User) error
	UserByPhonenumber(ctx context.Context, phonenumber string) (domain.User, error)
}
