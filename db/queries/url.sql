-- name: CreateUrl :one
INSERT INTO urls (token, url, created_at, expires_at, is_active)
VALUES (MD5($1), $1, NOW(), NOW() + INTERVAL '7 day', TRUE)
RETURNING token;

-- name: GetUrlByToken :one
SELECT Url, is_active FROM urls WHERE token = $1;

-- name: GetUrlTokenByUrl :one
SELECT token FROM urls WHERE url = $1;

-- name: CreateUrlWithExpire :one
INSERT INTO urls (token, url, created_at, expires_at, is_active)
VALUES (MD5($1), $1, NOW(), $2, TRUE)
RETURNING token;