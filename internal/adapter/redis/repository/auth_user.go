package repository

import (
	"context"

	"github.com/mar-cial/pizza-auth/internal/domain"
	"github.com/mar-cial/pizza-auth/internal/repository"
	"github.com/redis/go-redis/v9"
)

type authUsersRepo struct {
	client *redis.Client
}

func NewRedisAuthUsersRepo(client *redis.Client) repository.AuthUsers {
	return &authUsersRepo{client: client}
}

func (a *authUsersRepo) CreateUser(ctx context.Context, user domain.User) error {
	panic("not implemented") // TODO: Implement
}

func (a *authUsersRepo) UpdateUser(ctx context.Context, user domain.User) error {
	panic("not implemented") // TODO: Implement
}

func (a *authUsersRepo) UserByPhonenumber(ctx context.Context, phonenumber string) (domain.User, error) {
	panic("not implemented") // TODO: Implement
}
