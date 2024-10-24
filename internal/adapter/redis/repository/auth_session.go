package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mar-cial/pizza-auth/internal/repository"
	"github.com/redis/go-redis/v9"
)

var (
	ErrNotOk = errors.New("could not set session")
)

type authSessionRepo struct {
	client *redis.Client
}

func (a *authSessionRepo) CreateSession(ctx context.Context, userid, token string) error {
	_, err := a.client.Set(ctx, fmt.Sprintf("session:%s", userid), token, time.Hour*24*7).Result()
	if err != nil {
		return err
	}

	return nil
}

func (a *authSessionRepo) DeleteSession(ctx context.Context, userid string) error {
	sessionkey := fmt.Sprintf("session:%s", userid)

	_, err := a.client.Del(ctx, sessionkey).Result()

	return err
}

func NewRedisAuthSessionRepo(client *redis.Client) repository.AuthSession {
	return &authSessionRepo{client: client}
}
