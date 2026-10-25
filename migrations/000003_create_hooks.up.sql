CREATE TABLE hooks(
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY,
    hook_id UUID NOT NULL,
    hook_key TEXT NOT NULL,
    service TEXT NOT NULL,
    title TEXT NOT NULL,
    delivery_type TEXT NOT NULL,
    delivery_value TEXT NOT NULL,
    secret TEXT NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    total_called BIGINT,
    total_successful_delivery BIGINT,
    total_failed_delivery BIGINT,
    total_scheduled_delivery BIGINT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ,

    -- fk
    project_id UUID NOT NULL REFERENCES projects(project_id),

    -- pk
    PRIMARY KEY(hook_id)
);

CREATE INDEX idx_hook_project_id ON hooks(project_id);