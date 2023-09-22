-- name: CreateUser :one
INSERT INTO users (id, created_at, first_name, last_name, api_key)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING api_key;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;