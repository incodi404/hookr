CREATE TABLE delivery_history(
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    delivery_doc_id UUID NOT NULL,
    scheduled_at TIMESTAMPTZ NOT NULL,
    exponential_time_val INTERVAL NOT NULL DEFAULT INTERVAL '0',
    retry_counter INT DEFAULT 1,
    retry_type TEXT,
    status TEXT,
    error_data JSONB DEFAULT '{}'::jsonb,
    executed_at TIMESTAMPTZ,
    is_last_retry BOOLEAN DEFAULT FALSE,

    -- fk
    call_doc_id UUID NOT NULL REFERENCES hook_call_history(call_doc_id),

    -- pk
    PRIMARY KEY(delivery_doc_id)
);

CREATE INDEX idx_call_doc_id ON delivery_history(call_doc_id);