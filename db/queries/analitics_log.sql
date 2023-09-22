-- name: WriteLog :exec
INSERT INTO analitics_log (token_url, log_msg)
VALUES ($1, $2);

-- name: GetRedirectCount :one
SELECT COUNT(*) FROM analitics_log WHERE token_url = $1;