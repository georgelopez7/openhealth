package repository

import (
	"context"
	"database/sql"

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

// AddDoctorToAccountAssignment - assigns a doctor to an account
func (r *Repository) AddDoctorToAccountAssignment(ctx context.Context, assignment domain.DoctorAssignment) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		INSERT INTO doctor_to_account_assignments (id, account_id, doctor_id, valid_from, valid_to)
		VALUES ($1, $2, $3, $4, $5)
	`, assignment.ID, assignment.AccountID, assignment.DoctorID, assignment.ValidFrom, assignment.ValidTo)

	return err
}

// RemoveDoctorAssignment - removes a doctor assignment by setting valid_to to now
func (r *Repository) RemoveDoctorAssignment(ctx context.Context, id string) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		UPDATE doctor_to_account_assignments
		SET valid_to = GREATEST(NOW(), valid_from + '1 microsecond'::interval)
		WHERE id = $1
	`, id)

	return err
}

// RemoveAllDoctorAssignments - removes all open doctor assignments for a doctor
func (r *Repository) RemoveAllDoctorAssignments(ctx context.Context, doctorID string) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		UPDATE doctor_to_account_assignments
		SET valid_to = GREATEST(NOW(), valid_from + '1 microsecond'::interval)
		WHERE doctor_id = $1
		AND valid_to IS NULL
	`, doctorID)

	return err
}

// GetDoctorAssignmentByID - get a doctor assignment by its ID
func (r *Repository) GetDoctorAssignmentByID(ctx context.Context, id string) (*domain.DoctorAssignment, error) {
	var assignment domain.DoctorAssignment

	err := r.db.GetContext(ctx, &assignment, `
		SELECT id, account_id, doctor_id, valid_from, valid_to
		FROM doctor_to_account_assignments
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

// GetDoctorAssignments - get doctor assignments up to the provided limit
func (r *Repository) GetDoctorAssignments(ctx context.Context, limit int) ([]domain.DoctorAssignment, error) {
	var assignments []domain.DoctorAssignment

	err := r.db.SelectContext(ctx, &assignments, `
		SELECT id, account_id, doctor_id, valid_from, valid_to
		FROM doctor_to_account_assignments
		LIMIT $1
	`, limit)

	return assignments, err
}

// GetDoctorAssignmentsByDoctorID - get open doctor assignments for a doctor
func (r *Repository) GetDoctorAssignmentsByDoctorID(ctx context.Context, doctorID string) ([]domain.DoctorAssignment, error) {
	var assignments []domain.DoctorAssignment

	err := r.db.SelectContext(ctx, &assignments, `
		SELECT id, account_id, doctor_id, valid_from, valid_to
		FROM doctor_to_account_assignments
		WHERE doctor_id = $1
		AND valid_to IS NULL
	`, doctorID)

	return assignments, err
}
