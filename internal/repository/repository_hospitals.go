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
		TRUNCATE TABLE hospitals CASCADE
	`)

	return err
}

// AddHospitalAssignment - adds a hospital assignment to the database
func (r *Repository) AddHospitalAssignment(ctx context.Context, assignment domain.HospitalAssignment) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		INSERT INTO hospital_assignments (id, account_id, hospital_id, valid_from, valid_to)
		VALUES ($1, $2, $3, $4, $5)
	`, assignment.ID, assignment.AccountID, assignment.HospitalID, assignment.ValidFrom, assignment.ValidTo)

	return err
}

// RemoveHospitalAssignment - removes a hospital assignment by setting valid_to to now
func (r *Repository) RemoveHospitalAssignment(ctx context.Context, id string) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		UPDATE hospital_assignments
		SET valid_to = GREATEST(NOW(), valid_from + '1 microsecond'::interval)
		WHERE id = $1
	`, id)

	return err
}

// GetHospitalAssignmentByID - get a hospital assignment by its ID
func (r *Repository) GetHospitalAssignmentByID(ctx context.Context, id string) (*domain.HospitalAssignment, error) {
	var assignment domain.HospitalAssignment

	err := r.db.GetContext(ctx, &assignment, `
		SELECT id, account_id, hospital_id, valid_from, valid_to
		FROM hospital_assignments
		WHERE id = $1
	`, id)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &assignment, nil
	default:
		return nil, err
	}
}

// GetHospitalAssignments - get hospital assignments up to the provided limit
func (r *Repository) GetHospitalAssignments(ctx context.Context, limit int) ([]domain.HospitalAssignment, error) {
	var assignments []domain.HospitalAssignment

	err := r.db.SelectContext(ctx, &assignments, `
		SELECT id, account_id, hospital_id, valid_from, valid_to
		FROM hospital_assignments
		LIMIT $1
	`, limit)

	return assignments, err
}

// GetHospitalIDByAccountID - get the currently assigned hospital ID for an account
func (r *Repository) GetHospitalIDByAccountID(ctx context.Context, accountID string) (string, error) {
	var hospitalID string

	err := r.db.GetContext(ctx, &hospitalID, `
		SELECT hospital_id
		FROM hospital_assignments
		WHERE account_id = $1
		AND valid_to IS NULL
		ORDER BY valid_from DESC
		LIMIT 1
	`, accountID)

	switch err {
	case sql.ErrNoRows:
		return "", nil
	case nil:
		return hospitalID, nil
	default:
		return "", err
	}
}
