package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mar-cial/pizza-auth/internal/domain"
	"github.com/mar-cial/pizza-auth/internal/repository"
	"github.com/redis/go-redis/v9"
)

type lookupRepo struct {
	client *redis.Client
}

func (l *lookupRepo) UserByPhonenumber(ctx context.Context, phonenumber string) (*domain.User, error) {
	userStr, err := l.client.Get(ctx, fmt.Sprintf("user:phonenumber:%s", phonenumber)).Result()
	if err != nil {
		return nil, err
	}

	var user *domain.User
	if err := json.Unmarshal([]byte(userStr), &user); err != nil {
		return nil, err
	}

	return user, nil
}

func NewRedisLookupRepo(client *redis.Client) repository.AuthLookup {
	return &lookupRepo{client: client}
}
