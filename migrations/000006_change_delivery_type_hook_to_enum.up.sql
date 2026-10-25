CREATE TYPE hook_delivery_type AS ENUM (
    'EMAIL',
    'CUSTOM_API'
);

ALTER TABLE hooks
ALTER COLUMN delivery_type TYPE hook_delivery_type
USING delivery_type::hook_delivery_type;