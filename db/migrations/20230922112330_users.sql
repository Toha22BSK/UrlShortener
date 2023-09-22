-- migrate:up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT encode(sha256(random()::text::bytea), 'hex')
);

-- migrate:down

DROP TABLE users;