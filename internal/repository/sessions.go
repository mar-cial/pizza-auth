package repository

import (
	"context"
)

type AuthSession interface {
	CreateSession(ctx context.Context, userID, token string) error
	DeleteSession(ctx context.Context, sessionToken string) error
}
