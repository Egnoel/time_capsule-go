package repository

import (
	"context"

	"github.com/egnoel/future-message-go/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LetterRepository struct {
	db *pgxpool.Pool
}

func NewLetterRepository(db *pgxpool.Pool) *LetterRepository {
	return &LetterRepository{db: db}
}

func (r *LetterRepository) CreateLetter(ctx context.Context, letter *models.Letter, userID string) error {

	query := "INSERT INTO letters (user_id, subject, body, deliver_at) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(ctx, query, userID, letter.Subject, letter.Body, letter.DeliverAt)
	return err
}

func (r *LetterRepository) GetDeliveredLetters(
	ctx context.Context,
	userID string,
) ([]models.Letter, error) {

	query := `
		SELECT id, user_id, subject, body,
		       deliver_at, delivered, created_at
		FROM letters
		WHERE user_id = $1
		  AND delivered = TRUE
		ORDER BY deliver_at DESC
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var letters []models.Letter

	for rows.Next() {
		var letter models.Letter

		err := rows.Scan(
			&letter.ID,
			&letter.UserID,
			&letter.Subject,
			&letter.Body,
			&letter.DeliverAt,
			&letter.Delivered,
			&letter.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		letters = append(letters, letter)
	}

	return letters, nil
}

func (r *LetterRepository) GetPendingLetters(
	ctx context.Context,
	userID string,
) ([]models.Letter, error) {

	query := `
		SELECT id, user_id, subject, body,
		       deliver_at, delivered, created_at
		FROM letters
		WHERE user_id = $1
		  AND delivered = FALSE
		ORDER BY deliver_at ASC
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var letters []models.Letter

	for rows.Next() {
		var letter models.Letter

		err := rows.Scan(
			&letter.ID,
			&letter.UserID,
			&letter.Subject,
			&letter.Body,
			&letter.DeliverAt,
			&letter.Delivered,
			&letter.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		letters = append(letters, letter)
	}

	return letters, nil
}

func (r *LetterRepository) MarkLetterAsDelivered(ctx context.Context, letterID string, userID string) error {
	query := "UPDATE letters SET delivered = TRUE WHERE id = $1 AND user_id = $2 AND delivered = FALSE"
	_, err := r.db.Exec(ctx, query, letterID, userID)
	return err
}

func (r *LetterRepository) GetLetterByID(ctx context.Context, letterID string, userID string) (*models.Letter, error) {
	query := "SELECT id, user_id, subject, body, deliver_at, delivered, created_at FROM letters WHERE id = $1 AND user_id = $2"
	row := r.db.QueryRow(ctx, query, letterID, userID)
	var letter models.Letter
	err := row.Scan(&letter.ID, &letter.UserID, &letter.Subject, &letter.Body, &letter.DeliverAt, &letter.Delivered, &letter.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &letter, nil
}

func (r *LetterRepository) DeleteLetter(ctx context.Context, letterID string, userID string) error {
	query := "DELETE FROM letters WHERE id = $1 AND user_id = $2 AND delivered = FALSE"
	_, err := r.db.Exec(ctx, query, letterID, userID)
	return err
}

func (r *LetterRepository) UpdateLetter(ctx context.Context, letterID string, updatedLetter *models.Letter, userID string) error {
	query := "UPDATE letters SET subject = $1, body = $2, deliver_at = $3 WHERE id = $4 AND user_id = $5 AND delivered = FALSE"
	_, err := r.db.Exec(ctx, query, updatedLetter.Subject, updatedLetter.Body, updatedLetter.DeliverAt, letterID, userID)
	return err
}
