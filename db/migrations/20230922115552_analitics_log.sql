-- migrate:up
CREATE TABLE analitics_log (
    id BIGSERIAL PRIMARY KEY,
    token_url TEXT REFERENCES urls(token) ON DELETE CASCADE NOT NULL,
    log_msg TEXT NOT NULL DEFAULT 'redirected',
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- migrate:down

DROP TABLE analitics_log;
