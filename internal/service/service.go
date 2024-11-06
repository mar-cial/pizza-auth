package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mar-cial/pizza-auth/internal/domain"
	"github.com/mar-cial/pizza-auth/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPhonenumberMissing = errors.New("phonenumber missing")
	ErrPhonenumberExists  = errors.New("account exists")
)

type AuthService interface {
	Register(ctx context.Context, user *domain.User) error
	Login(ctx context.Context, creds *domain.Credentials) error
	Logout(ctx context.Context, userid string) error
}

// I believe these interfaces are only going to grow,
// so its a good idea to separate them from the beginning,
// I BELIEVE.
type authService struct {
	users    repository.AuthUsers
	sessions repository.AuthSession
	lookup   repository.AuthLookup
}

func (a *authService) Register(ctx context.Context, user *domain.User) error {
	user, err := a.lookup.UserByPhonenumber(ctx, user.Phonenumber)
	if err != nil {
		return err
	}

	if err := a.users.CreateUser(ctx, user); err != nil {
		return err
	}

	token := uuid.NewString()

	return a.sessions.CreateSession(ctx, user.ID, token)
}

func (a *authService) Login(ctx context.Context, creds *domain.Credentials) error {
	user, err := a.lookup.UserByPhonenumber(ctx, creds.Phonenumber)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return err
	}

	token := uuid.NewString()

	return a.sessions.CreateSession(ctx, user.ID, token)
}

func (a *authService) Logout(ctx context.Context, userid string) error {
	return a.sessions.DeleteSession(ctx, userid)
}

func NewAuthService(ur repository.AuthUsers, sr repository.AuthSession, l repository.AuthLookup) AuthService {
	return &authService{users: ur}
}
