package repository

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
)

func TestNewRedisAuthSessionRepo(t *testing.T) {
	db, _ := redismock.NewClientMock()

	authSessionRepo := NewRedisAuthSessionRepo(db)

	if authSessionRepo == nil {
		t.Errorf("auth session repo is nil")
	}
}

func TestCreateSession(t *testing.T) {

	db, mock := redismock.NewClientMock()

	authSessionRepo := NewRedisAuthSessionRepo(db)

	mock.ExpectSet(redismock.CustomMatch(`user:session:*`), time.Hour*24*7).SetVal("OK")

	if err := authSessionRepo.CreateSession(context.Background(), "1"); err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
