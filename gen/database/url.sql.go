// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: url.sql

package database

import (
	"context"
	"time"
)

const changeActiveStatus = `-- name: ChangeActiveStatus :exec
UPDATE urls SET is_active = FALSE 
WHERE expires_at < NOW() AND is_active = TRUE
`

func (q *Queries) ChangeActiveStatus(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, changeActiveStatus)
	return err
}

const createUrl = `-- name: CreateUrl :one
INSERT INTO urls (token, url, created_at, expires_at, is_active)
VALUES (MD5($1), $1, NOW(), NOW() + INTERVAL '7 day', TRUE)
RETURNING token
`

func (q *Queries) CreateUrl(ctx context.Context, url string) (string, error) {
	row := q.db.QueryRowContext(ctx, createUrl, url)
	var token string
	err := row.Scan(&token)
	return token, err
}

const createUrlWithExpire = `-- name: CreateUrlWithExpire :one
INSERT INTO urls (token, url, created_at, expires_at, is_active)
VALUES (MD5($1), $1, NOW(), $2, TRUE)
RETURNING token
`

type CreateUrlWithExpireParams struct {
	Url       string
	ExpiresAt time.Time
}

func (q *Queries) CreateUrlWithExpire(ctx context.Context, arg CreateUrlWithExpireParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createUrlWithExpire, arg.Url, arg.ExpiresAt)
	var token string
	err := row.Scan(&token)
	return token, err
}

const deleteUrlByToken = `-- name: DeleteUrlByToken :exec
UPDATE urls SET is_active = FALSE WHERE token = $1
`

func (q *Queries) DeleteUrlByToken(ctx context.Context, token string) error {
	_, err := q.db.ExecContext(ctx, deleteUrlByToken, token)
	return err
}

const getUrlByToken = `-- name: GetUrlByToken :one
SELECT Url, is_active FROM urls WHERE token = $1
`

type GetUrlByTokenRow struct {
	Url      string
	IsActive bool
}

func (q *Queries) GetUrlByToken(ctx context.Context, token string) (GetUrlByTokenRow, error) {
	row := q.db.QueryRowContext(ctx, getUrlByToken, token)
	var i GetUrlByTokenRow
	err := row.Scan(&i.Url, &i.IsActive)
	return i, err
}

const getUrlTokenByUrl = `-- name: GetUrlTokenByUrl :one
SELECT token FROM urls WHERE url = $1
`

func (q *Queries) GetUrlTokenByUrl(ctx context.Context, url string) (string, error) {
	row := q.db.QueryRowContext(ctx, getUrlTokenByUrl, url)
	var token string
	err := row.Scan(&token)
	return token, err
}

const setActiveTrue = `-- name: SetActiveTrue :one
UPDATE urls SET is_active = TRUE, expires_at = NOW() + INTERVAL '7 day' WHERE url = $1 RETURNING token
`

func (q *Queries) SetActiveTrue(ctx context.Context, url string) (string, error) {
	row := q.db.QueryRowContext(ctx, setActiveTrue, url)
	var token string
	err := row.Scan(&token)
	return token, err
}

const setActiveTrueExpire = `-- name: SetActiveTrueExpire :one
UPDATE urls SET is_active = TRUE, expires_at = $2 WHERE url = $1 RETURNING token
`

type SetActiveTrueExpireParams struct {
	Url       string
	ExpiresAt time.Time
}

func (q *Queries) SetActiveTrueExpire(ctx context.Context, arg SetActiveTrueExpireParams) (string, error) {
	row := q.db.QueryRowContext(ctx, setActiveTrueExpire, arg.Url, arg.ExpiresAt)
	var token string
	err := row.Scan(&token)
	return token, err
}
