package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/egnoel/future-message-go/internal/api/middleware"
	"github.com/egnoel/future-message-go/internal/models"
	"github.com/egnoel/future-message-go/internal/service"
	"github.com/go-chi/chi/v5"
)

type LetterHandler struct {
	letterService *service.LetterService
}

func NewLetterHandler(letterService *service.LetterService) *LetterHandler {
	return &LetterHandler{letterService: letterService}
}

// CreateLetter handles the creation of a new letter.
func (h *LetterHandler) CreateLetter(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Subject   string    `json:"subject"`
		Body      string    `json:"body"`
		DeliverAt time.Time `json:"deliver_at"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.Subject == "" || req.Body == "" || req.DeliverAt.IsZero() {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	userID := middleware.GetUserID(r.Context())
	letter := &models.Letter{
		Subject:   req.Subject,
		Body:      req.Body,
		DeliverAt: req.DeliverAt,
	}
	err := h.letterService.CreateLetter(r.Context(), letter, userID)
	if err != nil {
		http.Error(w, "Failed to create letter: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Letter created successfully",
	})
}

// GetDeliveredLetters handles fetching delivered letters for a user.
func (h *LetterHandler) GetDeliveredLetters(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	letters, err := h.letterService.GetDeliveredLetters(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to fetch delivered letters: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(letters)
}

// GetPendingLetters handles fetching pending letters for a user.
func (h *LetterHandler) GetPendingLetters(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	letters, err := h.letterService.GetPendingLetters(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to fetch pending letters: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(letters)
}

// GetLetterByID handles fetching a specific letter by its ID.
func (h *LetterHandler) GetLetterByID(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	letterID := chi.URLParam(r, "id")
	letter, err := h.letterService.GetLetterByID(r.Context(), letterID, userID)
	if err != nil {
		http.Error(w, "Failed to fetch letter: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(letter)
}

// MarkLetterAsDelivered handles marking a letter as delivered.
func (h *LetterHandler) MarkLetterAsDelivered(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	letterID := chi.URLParam(r, "id")
	err := h.letterService.MarkLetterAsDelivered(r.Context(), letterID, userID)
	if err != nil {
		http.Error(w, "Failed to mark letter as delivered: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Letter marked as delivered successfully",
	})
}

// UpdateLetter handles updating an existing letter.
func (h *LetterHandler) UpdateLetter(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	letterID := chi.URLParam(r, "id")
	var req struct {
		Subject   string    `json:"subject"`
		Body      string    `json:"body"`
		DeliverAt time.Time `json:"deliver_at"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.Subject == "" || req.Body == "" || req.DeliverAt.IsZero() {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	updatedLetter := &models.Letter{
		Subject:   req.Subject,
		Body:      req.Body,
		DeliverAt: req.DeliverAt,
	}

	err := h.letterService.UpdateLetter(r.Context(), letterID, updatedLetter, userID)
	if err != nil {
		http.Error(w, "Failed to update letter: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Letter updated successfully",
	})
}

// DeleteLetter handles deleting a letter.
func (h *LetterHandler) DeleteLetter(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	letterID := chi.URLParam(r, "id")
	err := h.letterService.DeleteLetter(r.Context(), letterID, userID)
	if err != nil {
		http.Error(w, "Failed to delete letter: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Letter deleted successfully",
	})
}
