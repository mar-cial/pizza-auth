package service

import (
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/mar-cial/pizza-auth/internal/adapter/redis/repository"
)

func TestNewAuthService(t *testing.T) {
	db, _ := redismock.NewClientMock()

	userRepo := repository.NewRedisAuthUsersRepo(db)
	sessionsRepo := repository.NewRedisAuthSessionRepo(db)
	lookupRepo := repository.NewRedisLookupRepo(db)

	service := NewAuthService(userRepo, sessionsRepo, lookupRepo)

	if service == nil {
		t.Errorf("service is nil")
	}
}
