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

-- name: ChangeActiveStatus :exec
UPDATE urls SET is_active = FALSE 
WHERE expires_at < NOW() AND is_active = TRUE;

-- name: SetActiveTrue :one
UPDATE urls SET is_active = TRUE, expires_at = NOW() + INTERVAL '7 day' WHERE url = $1 RETURNING token;

-- name: SetActiveTrueExpire :one
UPDATE urls SET is_active = TRUE, expires_at = $2 WHERE url = $1 RETURNING token;

-- name: DeleteUrlByToken :exec
UPDATE urls SET is_active = FALSE WHERE token = $1;
