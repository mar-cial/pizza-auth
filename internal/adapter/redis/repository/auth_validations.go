package repository

import (
	"context"

	"github.com/mar-cial/pizza-auth/internal/repository"
	"github.com/redis/go-redis/v9"
)

type authValidationsRepo struct {
	client *redis.Client
}

func NewRedisAuthValidationsRepo(client *redis.Client) repository.AuthValidations {
	return &authValidationsRepo{client: client}
}

func (a *authValidationsRepo) PhonenumberExists(ctx context.Context, phonenumber string) error {
	panic("not implemented") // TODO: Implement
}
