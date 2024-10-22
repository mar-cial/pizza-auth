package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/mar-cial/pizza-auth/internal/domain"
	"github.com/mar-cial/pizza-auth/internal/service"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *authHandler {
	return &authHandler{service: service}
}

func (a *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := domain.User{
		Email:       r.FormValue("email"),
		Username:    r.FormValue("username"),
		Password:    r.FormValue("password"),
		Phonenumber: r.FormValue("phonenumber"),
	}

	if err := a.service.Register(r.Context(), user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	creds := domain.Credentials{
		Phonenumber: r.FormValue("phonenumber"),
		Password:    r.FormValue("password"),
	}

	if err := a.service.Login(r.Context(), creds); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("pizza-auth")

	if errors.Is(err, http.ErrNoCookie) {
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	seshBytes, err := json.Marshal(cookie.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var sesh domain.Session
	if err := json.Unmarshal(seshBytes, &sesh); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the cookie exists, overwrite it with an expired cookie to remove it
	expiredCookie := &http.Cookie{
		Name:     cookie.Name,
		Value:    "",
		Expires:  time.Unix(0, 0), // Set the expiration date to a past time
		MaxAge:   -1,              // Set MaxAge to -1 to immediately remove it
		Path:     "/",             // Ensure the path matches the original cookie's path
		HttpOnly: true,            // Same settings as the original cookie
		Secure:   true,            // Same settings as the original cookie
	}

	if err := a.service.Logout(r.Context(), sesh.Token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, expiredCookie)
}
