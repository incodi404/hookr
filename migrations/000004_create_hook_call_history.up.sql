CREATE TABLE hook_call_history(
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY,
    call_doc_id UUID NOT NULL,
    is_called_successfully BOOLEAN DEFAULT FALSE,
    delivery_status TEXT NOT NULL,
    response_data JSONB NOT NULL DEFAULT '{}'::jsonb,
    total_retry INT DEFAULT 0,
    called_at TIMESTAMPTZ,
    delivered_at TIMESTAMPTZ,

    -- fk
    hook_id UUID NOT NULL REFERENCES hooks(hook_id),

    -- pk
    PRIMARY KEY(call_doc_id)
);

CREATE INDEX idx_hook_id ON hook_call_history(hook_id);