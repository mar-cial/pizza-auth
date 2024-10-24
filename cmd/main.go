package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mar-cial/pizza-auth/internal/adapter/redis/repository"
	"github.com/mar-cial/pizza-auth/internal/handler"
	"github.com/mar-cial/pizza-auth/internal/service"
	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_DB"),
	})

	usersRepo := repository.NewRedisAuthUsersRepo(client)
	sessionsRepo := repository.NewRedisAuthSessionRepo(client)
	lookupRepo := repository.NewRedisLookupRepo(client)

	srv := service.NewAuthService(usersRepo, sessionsRepo, lookupRepo)

	handler := handler.NewAuthHandler(srv)

	mux := http.NewServeMux()

	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/logout", handler.Logout)
	mux.HandleFunc("/register", handler.Register)

	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
