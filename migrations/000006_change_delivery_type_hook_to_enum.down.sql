ALTER TABLE hooks
ALTER COLUMN delivery_type TYPE TEXT
USING delivery_type::TEXT;

DROP TYPE hook_delivery_type;