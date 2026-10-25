CREATE TABLE projects (
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    project_id UUID NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    total_hook BIGINT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    -- fk
    user_id UUID NOT NULL REFERENCES users(user_id),

    -- pk
    PRIMARY KEY(project_id)
);

CREATE INDEX idx_project_user_id ON projects(user_id)