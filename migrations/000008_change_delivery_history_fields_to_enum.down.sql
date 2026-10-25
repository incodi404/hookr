ALTER TABLE delivery_history
ALTER COLUMN status DROP DEFAULT
ALTER COLUMN retry_type TYPE TEXT
USING retry_type::TEXT
ALTER COLUMN status TYPE TEXT
USING status::TEXT;

DROP TYPE hook_return_type_enum;
DROP TYPE hook_delivery_status_enum;