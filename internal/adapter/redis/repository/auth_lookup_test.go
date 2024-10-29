package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/mar-cial/pizza-auth/internal/domain"
	"github.com/redis/go-redis/v9"
)

func TestUserByPhonenumber(t *testing.T) {
	client, mock := redismock.NewClientMock()

	repo := NewRedisLookupRepo(client)

	ctx := context.Background()

	phonenumber := "1231231234"

	user := &domain.User{
		ID:          "1",
		Email:       "test@gmail.com",
		Username:    "username",
		Password:    "password",
		Phonenumber: phonenumber,
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	key := fmt.Sprintf("user:phonenumber:%s", user.Phonenumber)

	t.Run("Should return a user", func(t *testing.T) {
		mock.ExpectGet(key).SetVal(string(userBytes))

		userResult, err := repo.UserByPhonenumber(ctx, phonenumber)
		if err != nil {
			t.Error(err)
		}

		if userResult == nil {
			t.Error(err)
		}
	})

	t.Run("Should return an invalid user", func(t *testing.T) {
		mock.ExpectGet(key).SetErr(redis.Nil)

		userResult, err := repo.UserByPhonenumber(ctx, phonenumber)
		if err == nil {
			t.Errorf("did not return redis.Nil")
		}

		if userResult != nil {
			t.Errorf("user should be nil")
		}

	})
}
