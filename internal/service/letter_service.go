package service

import (
	"context"

	"github.com/egnoel/future-message-go/internal/models"
	"github.com/egnoel/future-message-go/internal/repository"
)


type LetterService struct {
	letterRepo *repository.LetterRepository
}

func NewLetterService(letterRepo *repository.LetterRepository) *LetterService {
	return &LetterService{letterRepo: letterRepo}
}

func (s *LetterService) CreateLetter(ctx context.Context, letter *models.Letter, userID string) error {
	return s.letterRepo.CreateLetter(ctx, letter, userID)
}

func (s *LetterService) GetDeliveredLetters(ctx context.Context, userID string) ([]models.Letter, error) {
	deliveredLetters, err := s.letterRepo.GetDeliveredLetters(ctx, userID)
	if err != nil {
		return nil, err
	}
	return deliveredLetters, nil
}

func (s *LetterService) GetPendingLetters(ctx context.Context, userID string) ([]models.Letter, error) {
	pendingLetters, err := s.letterRepo.GetPendingLetters(ctx, userID)
	if err != nil {
		return nil, err
	}
	return pendingLetters, nil
}

func (s *LetterService) GetLetterByID(ctx context.Context, letterID int64, userID string) (*models.Letter, error) {
	letter, err := s.letterRepo.GetLetterByID(ctx, letterID, userID)
	if err != nil {
		return nil, err
	}
	return letter, nil
}

func (s *LetterService) MarkLetterAsDelivered(ctx context.Context, letterID int64, userID string) error {
	return s.letterRepo.MarkLetterAsDelivered(ctx, letterID, userID)
}

func (s *LetterService) UpdateLetter(ctx context.Context, letterID int64, updatedLetter *models.Letter, userID string) error {
	return s.letterRepo.UpdateLetter(ctx, letterID, updatedLetter, userID)
}

func (s *LetterService) DeleteLetter(ctx context.Context, letterID int64, userID string) error {
	return s.letterRepo.DeleteLetter(ctx, letterID, userID)
}