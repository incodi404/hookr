-- name: CreateUser :one
INSERT INTO users (
    user_id,
    user_name,
    email_id,
    is_verified,
    created_at,
    verified_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

-- name: GetUserByUserId :one
SELECT * FROM users
WHERE user_id = $1
LIMIT 1;

-- name: GetUserByEmailId :one
SELECT * FROM users
WHERE email_id = $1
LIMIT 1;