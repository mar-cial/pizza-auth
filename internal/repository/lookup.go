package repository

import (
	"context"

	"github.com/mar-cial/pizza-auth/internal/domain"
)

type AuthLookup interface {
	UserByPhonenumber(ctx context.Context, phonenumber string) (*domain.User, error)
}
