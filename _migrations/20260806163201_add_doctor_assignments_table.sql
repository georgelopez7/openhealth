-- +goose Up
-- +goose StatementBegin
CREATE TABLE doctor_assignments (
    id TEXT PRIMARY KEY,
    account_id TEXT NOT NULL REFERENCES accounts(id),
    doctor_id TEXT NOT NULL REFERENCES doctors(id),
    valid_from TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    valid_to TIMESTAMPTZ,
    CHECK (valid_to IS NULL OR valid_to > valid_from)
);

CREATE INDEX idx_doctor_assignments_account ON doctor_assignments (account_id, valid_from);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS doctor_assignments;
-- +goose StatementEnd
