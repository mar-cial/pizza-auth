package service

import (
	"context"

	"github.com/mar-cial/pizza-auth/internal/domain"
	"github.com/mar-cial/pizza-auth/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, user domain.User) error
	Login(ctx context.Context, creds domain.Credentials) error
	Logout(ctx context.Context, sessionToken string) error
}

type authService struct {
	users    repository.AuthUsers
	sessions repository.AuthSession
}

func NewAuthService(usersRepo repository.AuthUsers, sessionsRepo repository.AuthSession) AuthService {
	return &authService{users: usersRepo, sessions: sessionsRepo}
}

func (a *authService) Register(ctx context.Context, user domain.User) error {
	user, err := a.users.UserByPhonenumber(ctx, user.Phonenumber)
	if err != nil {
		return err
	}

	if err := a.users.CreateUser(ctx, user); err != nil {
		return err
	}

	return a.sessions.CreateSession(ctx, user.ID)
}

func (a *authService) Login(ctx context.Context, creds domain.Credentials) error {
	user, err := a.users.UserByPhonenumber(ctx, creds.Phonenumber)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return err
	}

	return a.sessions.CreateSession(ctx, user.ID)
}

func (a *authService) Logout(ctx context.Context, sessionToken string) error {
	return a.sessions.DeleteSession(ctx, sessionToken)
}
