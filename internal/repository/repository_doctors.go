package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"openhealth/internal/domain"
	"openhealth/internal/pkg/postgres"
)

// AddDoctor - adds a new doctor to the database
func (r *Repository) AddDoctor(ctx context.Context, doctor domain.Doctor) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		INSERT INTO doctors (id, account_id, status, created_at)
		VALUES ($1, $2, $3, $4)
	`, doctor.ID, doctor.AccountID, doctor.Status, doctor.CreatedAt)

	return err
}

// GetDoctorByID - get a doctor by its ID
func (r *Repository) GetDoctorByID(ctx context.Context, id string) (*domain.Doctor, error) {
	var doctor domain.Doctor

	err := r.db.GetContext(ctx, &doctor, `
		SELECT id, account_id, status, created_at
		FROM doctors
		WHERE id = $1
	`, id)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &doctor, nil
	default:
		return nil, err
	}
}

// GetDoctors - get doctors up to the provided limit
func (r *Repository) GetDoctors(ctx context.Context, limit int) ([]domain.Doctor, error) {
	var doctors []domain.Doctor

	err := r.db.SelectContext(ctx, &doctors, `
		SELECT id, account_id, status, created_at
		FROM doctors
		LIMIT $1
	`, limit)

	return doctors, err
}

// UpdateDoctorStatusByID - updates the status of an existing doctor
func (r *Repository) UpdateDoctorStatusByID(ctx context.Context, id string, status domain.DoctorStatus) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		UPDATE doctors
		SET status = $1
		WHERE id = $2
	`, status, id)

	return err
}

// AssignDoctorToAccount - assigns a doctor to an account
func (r *Repository) AssignDoctorToAccount(ctx context.Context, accountID string, doctorID string) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		INSERT INTO doctor_assignments (id, account_id, doctor_id)
		VALUES ($1, $2, $3)
	`, uuid.NewString(), accountID, doctorID)

	return err
}

// ExpireDoctorAssignments - expires all open assignments for a doctor
func (r *Repository) ExpireDoctorAssignments(ctx context.Context, doctorID string) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		UPDATE doctor_assignments
		SET valid_to = NOW()
		WHERE doctor_id = $1
		AND valid_to IS NULL
	`, doctorID)

	return err
}
