package repository

import (
	"context"
	"database/sql"

	"openhealth/internal/domain"
	"openhealth/internal/pkg/postgres"
)

// AddHospital - adds a new hospital to the database
func (r *Repository) AddHospital(ctx context.Context, hospital domain.Hospital) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		INSERT INTO hospitals (id, name, status, created_at)
		VALUES ($1, $2, $3, $4)
	`, hospital.ID, hospital.Name, hospital.Status, hospital.CreatedAt)

	return err
}

// GetHospitalByID - get a hospital by its ID
func (r *Repository) GetHospitalByID(ctx context.Context, id string) (*domain.Hospital, error) {
	var hospital domain.Hospital

	err := r.db.GetContext(ctx, &hospital, `
		SELECT id, name, status, created_at
		FROM hospitals
		WHERE id = $1
	`, id)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &hospital, nil
	default:
		return nil, err
	}
}

// GetHospitals - get hospitals up to the provided limit
func (r *Repository) GetHospitals(ctx context.Context, limit int) ([]domain.Hospital, error) {
	var hospitals []domain.Hospital

	err := r.db.SelectContext(ctx, &hospitals, `
		SELECT id, name, status, created_at
		FROM hospitals
		LIMIT $1
	`, limit)

	return hospitals, err
}

// UpdateHospital - updates an existing hospital
func (r *Repository) UpdateHospital(ctx context.Context, hospital domain.Hospital) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		UPDATE hospitals
		SET name = $1, status = $2
		WHERE id = $3
	`, hospital.Name, hospital.Status, hospital.ID)

	return err
}

// ArchiveHospital - archives an existing hospital
func (r *Repository) ArchiveHospital(ctx context.Context, id string) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		UPDATE hospitals
		SET status = $1
		WHERE id = $2
	`, domain.HospitalStatusArchived, id)

	return err
}

// ResetHospitals - resets all hospitals
func (r *Repository) ResetHospitals(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
		TRUNCATE TABLE hospitals
	`)

	return err
}
