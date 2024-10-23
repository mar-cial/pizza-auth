package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mar-cial/pizza-auth/internal/domain"
	"github.com/mar-cial/pizza-auth/internal/repository"
	"github.com/redis/go-redis/v9"
)

var (
	ErrNotOk = errors.New("could not set session")
)

type authSessionRepo struct {
	client *redis.Client
}

func NewRedisAuthSessionRepo(client *redis.Client) repository.AuthSession {
	return &authSessionRepo{client: client}
}

func (a *authSessionRepo) CreateSession(ctx context.Context, userID, token string) error {
	session := domain.Session{
		Token:  token,
		UserID: userID,
	}

	sessionkey := fmt.Sprintf("user:session:%s", token)

	result, err := a.client.Set(ctx, sessionkey, session.UserID, time.Hour*24*7).Result()
	if err != nil {
		return err
	}

	if result != "OK" {
		return ErrNotOk
	}

	return nil
}

func (a *authSessionRepo) DeleteSession(ctx context.Context, sessionToken string) error {
	sessionkey := fmt.Sprintf("user:session:%s", sessionToken)

	_, err := a.client.Del(ctx, sessionkey).Result()

	return err
}
