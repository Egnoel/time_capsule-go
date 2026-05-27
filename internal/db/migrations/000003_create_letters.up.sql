CREATE TABLE letters (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id      UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    subject      TEXT NOT NULL,
    body         TEXT NOT NULL,
    deliver_at   TIMESTAMPTZ NOT NULL,
    delivered    BOOLEAN NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_letters_deliver_at ON letters(deliver_at) WHERE delivered = FALSE;