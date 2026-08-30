package repository

import (
	"context"
	"database/sql"

	"openhealth/internal/domain"
	"openhealth/internal/pkg/postgres"
)

// AddMedicalRecord - adds a new medical record to the database
func (r *Repository) AddMedicalRecord(ctx context.Context, record domain.MedicalRecord) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		INSERT INTO medical_records (id, account_id, title, description, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, record.ID, record.AccountID, record.Title, record.Description, record.Status, record.CreatedAt, record.UpdatedAt)

	return err
}

// GetMedicalRecordByID - get a medical record by its ID
func (r *Repository) GetMedicalRecordByID(ctx context.Context, id string) (*domain.MedicalRecord, error) {
	var record domain.MedicalRecord

	err := r.db.GetContext(ctx, &record, `
		SELECT id, account_id, title, description, status, created_at, updated_at
		FROM medical_records
		WHERE id = $1
	`, id)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &record, nil
	default:
		return nil, err
	}
}

// UpdateMedicalRecord - updates an existing medical record
func (r *Repository) UpdateMedicalRecord(ctx context.Context, record domain.MedicalRecord) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		UPDATE medical_records
		SET title = $1, description = $2, status = $3, updated_at = $4
		WHERE id = $5
	`, record.Title, record.Description, record.Status, record.UpdatedAt, record.ID)

	return err
}

// ArchiveMedicalRecord - archives an existing medical record
func (r *Repository) ArchiveMedicalRecord(ctx context.Context, id string) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		UPDATE medical_records
		SET status = $1, updated_at = NOW()
		WHERE id = $2
	`, domain.MedicalRecordStatusArchived, id)

	return err
}

// GetMedicalRecords - get all medical records
func (r *Repository) GetMedicalRecords(ctx context.Context) ([]domain.MedicalRecord, error) {
	var records = make([]domain.MedicalRecord, 0)

	err := r.db.SelectContext(ctx, &records, `
		SELECT id, account_id, title, description, status, created_at, updated_at
		FROM medical_records
		ORDER BY created_at DESC
	`)

	return records, err
}

// GetMedicalRecordsByAccountID - get medical records for an account
func (r *Repository) GetMedicalRecordsByAccountID(ctx context.Context, accountID string) ([]domain.MedicalRecord, error) {
	var records = make([]domain.MedicalRecord, 0)

	err := r.db.SelectContext(ctx, &records, `
		SELECT id, account_id, title, description, status, created_at, updated_at
		FROM medical_records
		WHERE account_id = $1
		ORDER BY created_at DESC
	`, accountID)

	return records, err
}

// ResetMedicalRecords - resets all medical records
func (r *Repository) ResetMedicalRecords(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
		TRUNCATE TABLE medical_records CASCADE
	`)

	return err
}
