package repository

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
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

	mock.Regexp().ExpectHSet(`^user:[a-z0-9\-]+$`, user).SetVal(1)
	if err := authUsersRepo.CreateUser(ctx, user); err != nil {
		t.Error(err)
	}
}
