package repository

import (
	"context"

	"github.com/mar-cial/pizza-auth/internal/domain"
)

type AuthUsers interface {
	CreateUser(ctx context.Context, user *domain.User) error
}
