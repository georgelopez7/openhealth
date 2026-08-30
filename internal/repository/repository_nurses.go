package repository

import (
	"context"
	"database/sql"

	"openhealth/internal/domain"
	"openhealth/internal/pkg/postgres"
)

// AddNurse - adds a new nurse to the database
func (r *Repository) AddNurse(ctx context.Context, nurse domain.Nurse) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		INSERT INTO nurses (id, account_id, status, created_at)
		VALUES ($1, $2, $3, $4)
	`, nurse.ID, nurse.AccountID, nurse.Status, nurse.CreatedAt)

	return err
}

// GetNurseByID - get a nurse by its ID
func (r *Repository) GetNurseByID(ctx context.Context, id string) (*domain.Nurse, error) {
	var nurse domain.Nurse

	err := r.db.GetContext(ctx, &nurse, `
		SELECT id, account_id, status, created_at
		FROM nurses
		WHERE id = $1
	`, id)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &nurse, nil
	default:
		return nil, err
	}
}

// GetNurseByAccountID - get a nurse by its associated account ID
func (r *Repository) GetNurseByAccountID(ctx context.Context, accountID string) (*domain.Nurse, error) {
	var nurse domain.Nurse

	err := r.db.GetContext(ctx, &nurse, `
		SELECT id, account_id, status, created_at
		FROM nurses
		WHERE account_id = $1
	`, accountID)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &nurse, nil
	default:
		return nil, err
	}
}

// GetNurses - get nurses up to the provided limit
func (r *Repository) GetNurses(ctx context.Context, limit int) ([]domain.Nurse, error) {
	var nurses = make([]domain.Nurse, 0)

	err := r.db.SelectContext(ctx, &nurses, `
		SELECT id, account_id, status, created_at
		FROM nurses
		LIMIT $1
	`, limit)

	return nurses, err
}

// UpdateNurseStatusByID - updates the status of an existing nurse
func (r *Repository) UpdateNurseStatusByID(ctx context.Context, id string, status domain.NurseStatus) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		UPDATE nurses
		SET status = $1
		WHERE id = $2
	`, status, id)

	return err
}

// AddNurseToHospitalAssignment - adds a nurse to hospital assignment to the database
func (r *Repository) AddNurseToHospitalAssignment(ctx context.Context, assignment domain.NurseToHospitalAssignment) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		INSERT INTO nurse_to_hospital_assignments (id, nurse_id, hospital_id, valid_from, valid_to)
		VALUES ($1, $2, $3, $4, $5)
	`, assignment.ID, assignment.NurseID, assignment.HospitalID, assignment.ValidFrom, assignment.ValidTo)

	return err
}

// RemoveNurseToHospitalAssignment - removes a nurse to hospital assignment by setting valid_to to now
func (r *Repository) RemoveNurseToHospitalAssignment(ctx context.Context, id string) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		UPDATE nurse_to_hospital_assignments
		SET valid_to = GREATEST(NOW(), valid_from + '1 microsecond'::interval)
		WHERE id = $1
	`, id)

	return err
}

// RemoveAllNurseToHospitalAssignments - removes all open nurse to hospital assignments for a nurse
func (r *Repository) RemoveAllNurseToHospitalAssignments(ctx context.Context, nurseID string) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		UPDATE nurse_to_hospital_assignments
		SET valid_to = GREATEST(NOW(), valid_from + '1 microsecond'::interval)
		WHERE nurse_id = $1
		AND valid_to IS NULL
	`, nurseID)

	return err
}

// GetNurseToHospitalAssignmentByID - get a nurse to hospital assignment by its ID
func (r *Repository) GetNurseToHospitalAssignmentByID(ctx context.Context, id string) (*domain.NurseToHospitalAssignment, error) {
	var assignment domain.NurseToHospitalAssignment

	err := r.db.GetContext(ctx, &assignment, `
		SELECT id, nurse_id, hospital_id, valid_from, valid_to
		FROM nurse_to_hospital_assignments
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

// GetNurseToHospitalAssignments - get nurse to hospital assignments up to the provided limit
func (r *Repository) GetNurseToHospitalAssignments(ctx context.Context, limit int) ([]domain.NurseToHospitalAssignment, error) {
	var assignments = make([]domain.NurseToHospitalAssignment, 0)

	err := r.db.SelectContext(ctx, &assignments, `
		SELECT id, nurse_id, hospital_id, valid_from, valid_to
		FROM nurse_to_hospital_assignments
		LIMIT $1
	`, limit)

	return assignments, err
}

// GetNurseToHospitalAssignmentsByNurseID - get open nurse to hospital assignments for a nurse
func (r *Repository) GetNurseToHospitalAssignmentsByNurseID(ctx context.Context, nurseID string) ([]domain.NurseToHospitalAssignment, error) {
	var assignments = make([]domain.NurseToHospitalAssignment, 0)

	err := r.db.SelectContext(ctx, &assignments, `
		SELECT id, nurse_id, hospital_id, valid_from, valid_to
		FROM nurse_to_hospital_assignments
		WHERE nurse_id = $1
		AND valid_to IS NULL
	`, nurseID)

	return assignments, err
}

// GetHospitalIDByNurseID - get the currently assigned hospitalID for a nurse
func (r *Repository) GetHospitalIDByNurseID(ctx context.Context, nurseID string) (string, error) {
	var hospitalID string

	err := r.db.GetContext(ctx, &hospitalID, `
		SELECT hospital_id
		FROM nurse_to_hospital_assignments
		WHERE nurse_id = $1
		AND valid_to IS NULL
		ORDER BY valid_from DESC
		LIMIT 1
	`, nurseID)

	switch err {
	case sql.ErrNoRows:
		return "", nil
	case nil:
		return hospitalID, nil
	default:
		return "", err
	}
}

// GetNursesByHospitalID - get open nurse assignments for a hospital
func (r *Repository) GetNursesByHospitalID(ctx context.Context, hospitalID string) ([]domain.NurseToHospitalAssignment, error) {
	var assignments = make([]domain.NurseToHospitalAssignment, 0)

	err := r.db.SelectContext(ctx, &assignments, `
		SELECT id, nurse_id, hospital_id, valid_from, valid_to
		FROM nurse_to_hospital_assignments
		WHERE hospital_id = $1
		AND valid_to IS NULL
	`, hospitalID)

	return assignments, err
}
