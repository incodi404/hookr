CREATE TYPE hook_status_enum AS ENUM(
    'SUCCESS',
    'FAILED',
    'RETRYING',
    'PENDING'
);

ALTER TABLE hook_call_history
ALTER COLUMN delivery_status TYPE hook_status_enum
USING delivery_status::hook_status_enum;

ALTER TABLE hook_call_history
ALTER COLUMN delivery_status SET DEFAULT 'PENDING';