package repository

import (
	"context"
)

type AuthSession interface {
	CreateSession(ctx context.Context, userid, token string) error
	DeleteSession(ctx context.Context, userid string) error
}
