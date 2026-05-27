package models

import (
	"time"

	"github.com/google/uuid"
)

type Letter struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	Subject   string    `db:"subject"`
	Body      string    `db:"body"`
	DeliverAt time.Time `db:"deliver_at"`
	Delivered bool      `db:"delivered"`
	CreatedAt time.Time `db:"created_at"`
}
