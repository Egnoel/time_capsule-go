package service

import (
	"context"
	"errors"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/egnoel/future-message-go/internal/models"
	"github.com/egnoel/future-message-go/internal/repository"
	"github.com/egnoel/future-message-go/pkg/password"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type AuthService struct {
	userRepo       *repository.UserRepository
	sessionManager *scs.SessionManager
}

func NewAuthService(userRepo *repository.UserRepository, sessionManager *scs.SessionManager) *AuthService {
	return &AuthService{
		userRepo:       userRepo,
		sessionManager: sessionManager,
	}
}

func (s *AuthService) Register(ctx context.Context, email, name, rawPassword string) error {
	hashedPassword, err := password.HashPassword(rawPassword)
	if err != nil {
		log.Println("Error hashing password:", err)
		return err
	}

	user := &models.User{
		Email:    email,
		Name:     name,
		Password: hashedPassword,
	}

	return s.userRepo.CreateUser(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, email, rawPassword string) error {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return ErrInvalidCredentials
	}

	if !password.CheckPasswordHash(rawPassword, user.Password) {
		return ErrInvalidCredentials
	}

	s.sessionManager.Put(ctx, "user_id", user.ID.String())
	s.sessionManager.Put(ctx, "user_email", user.Email)

	return nil
}

func (s *AuthService) Logout(ctx context.Context) error {
	return s.sessionManager.Destroy(ctx)
}
