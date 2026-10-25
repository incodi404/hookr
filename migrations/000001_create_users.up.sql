CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    user_id UUID NOT NULL UNIQUE,
    user_name TEXT,
    email_id TEXT NOT NULL UNIQUE,
    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    is_blocked BOOLEAN NOT NULL DEFAULT FALSE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    verified_at TIMESTAMPTZ,

    PRIMARY KEY(user_id) -- primary key
);

-- Indexes
CREATE INDEX idx_email_id ON users (email_id);