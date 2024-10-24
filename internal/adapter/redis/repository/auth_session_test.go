package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/google/uuid"
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

	ctx := context.Background()

	userid := uuid.NewString()
	token := uuid.NewString()
	key := fmt.Sprintf("session:%s", userid)

	mock.ExpectSet(key, token, time.Hour*24*7).SetVal("OK")
	if err := authSessionRepo.CreateSession(ctx, userid, token); err != nil {
		t.Error(err)
	}

	mock.ExpectDel(key).SetErr(fmt.Errorf("redis error"))
	if err := authSessionRepo.DeleteSession(ctx, userid); err == nil {
		t.Error("expected an error but got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestDeleteSession(t *testing.T) {
	db, mock := redismock.NewClientMock()

	authSessionRepo := NewRedisAuthSessionRepo(db)

	ctx := context.Background()
	userid := uuid.NewString()
	key := fmt.Sprintf("session:%s", userid)

	mock.ExpectDel(key).SetVal(1)
	if err := authSessionRepo.DeleteSession(ctx, userid); err != nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
