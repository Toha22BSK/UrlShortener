-- migrate:up
CREATE TABLE urls (
    token TEXT PRIMARY KEY,
    url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL DEFAULT NOW() + INTERVAL '7 day',
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

-- migrate:down

DROP TABLE urls;