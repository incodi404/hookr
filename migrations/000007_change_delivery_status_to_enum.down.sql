ALTER TABLE hook_call_history
ALTER COLUMN delivery_status DROP DEFAULT
ALTER COLUMN delivery_status TYPE TEXT
USING delivery_status::TEXT;

DROP TYPE hook_delivery_status_enum;