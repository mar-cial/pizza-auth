package repository

import (
	"context"
)

type AuthSession interface {
	CreateSession(ctx context.Context, userID string) error
	DeleteSession(ctx context.Context, sessionToken string) error
}
