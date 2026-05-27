package main

import (
	"log"
	"net/http"
	"os"
	 "github.com/joho/godotenv"

	"github.com/egnoel/future-message-go/internal/api"
	"github.com/egnoel/future-message-go/internal/api/handlers"
	"github.com/egnoel/future-message-go/internal/repository"
	"github.com/egnoel/future-message-go/internal/service"
	"github.com/egnoel/future-message-go/pkg/database"
	sessionpkg "github.com/egnoel/future-message-go/pkg/session"
)

func main() {

	_ = godotenv.Load()

	// =====================================
	// DATABASE
	// =====================================

	dbPool, err := database.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	// =====================================
	// REDIS / SESSION
	// =====================================

	redisPool := sessionpkg.NewRedisPool()

	sessionManager := sessionpkg.NewManager(redisPool)

	// Em desenvolvimento local
	sessionManager.Cookie.Secure = false

	// =====================================
	// REPOSITORIES
	// =====================================

	userRepo := repository.NewUserRepository(dbPool)

	// =====================================
	// SERVICES
	// =====================================

	authService := service.NewAuthService(
		userRepo,
		sessionManager,
	)

	// =====================================
	// HANDLERS
	// =====================================

	authHandler := handlers.NewAuthHandler(authService)

	// =====================================
	// ROUTER
	// =====================================

	router := api.NewRouter(
		authHandler,
		sessionManager,
	)

	// =====================================
	// SERVER
	// =====================================

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server running on port %s", port)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
