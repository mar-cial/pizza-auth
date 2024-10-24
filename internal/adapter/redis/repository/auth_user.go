package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mar-cial/pizza-auth/internal/domain"
	"github.com/mar-cial/pizza-auth/internal/repository"
	"github.com/redis/go-redis/v9"
)

type authUsersRepo struct {
	client *redis.Client
}

func (a *authUsersRepo) CreateUser(ctx context.Context, user *domain.User) error {
	id := uuid.NewString()
	key := fmt.Sprintf("user:%s", id)

	_, err := a.client.HSet(ctx, key, user).Result()

	return err
}

func NewRedisAuthUsersRepo(client *redis.Client) repository.AuthUsers {
	return &authUsersRepo{client: client}
}
