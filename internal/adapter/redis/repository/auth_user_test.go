package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/google/uuid"
	"github.com/mar-cial/pizza-auth/internal/domain"
)

func TestNewRedisAuthUsersRepo(t *testing.T) {
	db, mock := redismock.NewClientMock()

	authUsersRepo := NewRedisAuthUsersRepo(db)

	ctx := context.Background()

	user := &domain.User{
		Email:       "test@gmail.com",
		Username:    "test",
		Password:    "test",
		Phonenumber: "1231231234",
	}

	userid := uuid.NewString()

	key := fmt.Sprintf("user:%s", userid)
	mock.ExpectHSet(key, user)
	if err := authUsersRepo.CreateUser(ctx, user); err != nil {
		t.Error(err)
	}
}
