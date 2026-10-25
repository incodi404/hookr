CREATE TYPE hook_return_type_enum AS ENUM ('AUTO', 'REPLAY');
CREATE TYPE hook_delivery_status_enum AS ENUM ('SUCCESS', 'FAILED', 'RETRYING', 'PENDING');

ALTER TABLE delivery_history
    ALTER COLUMN status DROP DEFAULT;

ALTER TABLE delivery_history
    ALTER COLUMN retry_type TYPE hook_return_type_enum
        USING retry_type::hook_return_type_enum,
    ALTER COLUMN status TYPE hook_delivery_status_enum
        USING status::hook_delivery_status_enum;

ALTER TABLE delivery_history
    ALTER COLUMN status SET DEFAULT 'PENDING';