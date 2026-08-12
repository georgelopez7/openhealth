-- +goose Up
-- +goose StatementBegin
CREATE TABLE nurses (
    id TEXT PRIMARY KEY,
    account_id TEXT NOT NULL UNIQUE,
    status TEXT NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (account_id) REFERENCES accounts(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS nurses;
-- +goose StatementEnd
