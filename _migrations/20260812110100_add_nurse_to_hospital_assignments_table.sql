-- +goose Up
-- +goose StatementBegin
CREATE TABLE nurse_to_hospital_assignments (
    id TEXT PRIMARY KEY,
    nurse_id TEXT NOT NULL REFERENCES nurses(id),
    hospital_id TEXT NOT NULL REFERENCES hospitals(id),
    valid_from TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    valid_to TIMESTAMPTZ,
    CHECK (valid_to IS NULL OR valid_to > valid_from)
);

CREATE INDEX idx_nurse_to_hospital_assignments_nurse ON nurse_to_hospital_assignments (nurse_id, valid_from);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS nurse_to_hospital_assignments;
-- +goose StatementEnd
