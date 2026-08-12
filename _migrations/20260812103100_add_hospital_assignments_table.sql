-- +goose Up
-- +goose StatementBegin
CREATE TABLE hospital_assignments (
    id TEXT PRIMARY KEY,
    account_id TEXT NOT NULL REFERENCES accounts(id),
    hospital_id TEXT NOT NULL REFERENCES hospitals(id),
    valid_from TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    valid_to TIMESTAMPTZ,
    CHECK (valid_to IS NULL OR valid_to > valid_from)
);

CREATE INDEX idx_hospital_assignments_account ON hospital_assignments (account_id, valid_from);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS hospital_assignments;
-- +goose StatementEnd
