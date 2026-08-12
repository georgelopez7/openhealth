package test

import (
	"database/sql"
	"openhealth/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRepository_AddDoctor(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	expected := newTestDoctor(t)

	err := repo.AddDoctor(ctx, *expected)
	require.NoError(t, err)

	doctor, err := repo.GetDoctorByID(ctx, expected.ID)
	require.NoError(t, err)
	require.Equal(t, expected.ID, doctor.ID)
	require.Equal(t, expected.AccountID, doctor.AccountID)
	require.Equal(t, domain.DoctorStatusActive, doctor.Status)
}

func TestRepository_GetDoctorByID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully return doctor", func(t *testing.T) {
		expected := newTestDoctor(t)

		err := repo.AddDoctor(ctx, *expected)
		require.NoError(t, err)

		doctor, err := repo.GetDoctorByID(ctx, expected.ID)
		require.NoError(t, err)
		require.Equal(t, expected.ID, doctor.ID)
		require.Equal(t, expected.AccountID, doctor.AccountID)
		require.Equal(t, domain.DoctorStatusActive, doctor.Status)
	})

	t.Run("should handle case when doctor is not found", func(t *testing.T) {
		doctor, err := repo.GetDoctorByID(ctx, "non-existent-id")
		require.NoError(t, err)
		require.Nil(t, doctor)
	})
}

func TestRepository_GetDoctors(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should return doctors up to the limit", func(t *testing.T) {
		drJones := newTestDoctor(t)
		drSmith := newTestDoctor(t)
		drBrown := newTestDoctor(t)

		require.NoError(t, repo.AddDoctor(ctx, *drJones))
		require.NoError(t, repo.AddDoctor(ctx, *drSmith))
		require.NoError(t, repo.AddDoctor(ctx, *drBrown))

		doctors, err := repo.GetDoctors(ctx, 2)
		require.NoError(t, err)
		require.Len(t, doctors, 2)

		createdIDs := []string{drJones.ID, drSmith.ID, drBrown.ID}

		require.Contains(t, createdIDs, doctors[0].ID)
		require.Contains(t, createdIDs, doctors[1].ID)
	})

	t.Run("should return all doctors when limit is greater than count", func(t *testing.T) {
		repo.ResetAccounts(ctx)

		expected := newTestDoctor(t)
		require.NoError(t, repo.AddDoctor(ctx, *expected))

		doctors, err := repo.GetDoctors(ctx, 10)
		require.NoError(t, err)
		require.Len(t, doctors, 1)
		require.Equal(t, expected.ID, doctors[0].ID)
	})

	t.Run("should return empty slice when no doctors exist", func(t *testing.T) {
		repo.ResetAccounts(ctx)

		doctors, err := repo.GetDoctors(ctx, 5)
		require.NoError(t, err)
		require.Empty(t, doctors)
	})
}

func TestRepository_UpdateDoctorStatusByID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully update doctor status", func(t *testing.T) {
		doctor := newTestDoctor(t)

		err := repo.AddDoctor(ctx, *doctor)
		require.NoError(t, err)

		err = repo.UpdateDoctorStatusByID(ctx, doctor.ID, domain.DoctorStatusArchived)
		require.NoError(t, err)

		archived, err := repo.GetDoctorByID(ctx, doctor.ID)
		require.NoError(t, err)
		require.Equal(t, domain.DoctorStatusArchived, archived.Status)
	})

	t.Run("should handle case when doctor is not found", func(t *testing.T) {
		err := repo.UpdateDoctorStatusByID(ctx, "non-existent-id", domain.DoctorStatusArchived)
		require.NoError(t, err)
	})
}

func TestRepository_AssignDoctorToAccount(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully assign a doctor to an account", func(t *testing.T) {
		doctor := newTestDoctor(t)
		require.NoError(t, repo.AddDoctor(ctx, *doctor))

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
		require.NoError(t, repo.AddAccount(ctx, *account))

		err := repo.AssignDoctorToAccount(ctx, account.ID, doctor.ID)
		require.NoError(t, err)

		var count int
		err = db.QueryRowContext(ctx, `
			SELECT COUNT(*) FROM doctor_assignments
			WHERE account_id = $1 AND doctor_id = $2 AND valid_to IS NULL
		`, account.ID, doctor.ID).Scan(&count)
		require.NoError(t, err)
		require.Equal(t, 1, count)
	})
}

func TestRepository_ExpireDoctorAssignments(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should expire all open assignments for a doctor", func(t *testing.T) {
		doctor := newTestDoctor(t)
		require.NoError(t, repo.AddDoctor(ctx, *doctor))

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
		require.NoError(t, repo.AddAccount(ctx, *account))

		_, err := db.ExecContext(ctx, `
			INSERT INTO doctor_assignments (id, account_id, doctor_id)
			VALUES ($1, $2, $3)
		`, "assignment-id-1", account.ID, doctor.ID)
		require.NoError(t, err)

		err = repo.ExpireDoctorAssignments(ctx, doctor.ID)
		require.NoError(t, err)

		var validTo sql.NullTime
		err = db.QueryRowContext(ctx, `
			SELECT valid_to FROM doctor_assignments
			WHERE id = $1
		`, "assignment-id-1").Scan(&validTo)
		require.NoError(t, err)
		require.True(t, validTo.Valid)
	})

	t.Run("should only expire open assignments", func(t *testing.T) {
		doctor := newTestDoctor(t)
		require.NoError(t, repo.AddDoctor(ctx, *doctor))

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
		require.NoError(t, repo.AddAccount(ctx, *account))

		now := time.Now().UTC()
		_, err := db.ExecContext(ctx, `
			INSERT INTO doctor_assignments (id, account_id, doctor_id, valid_to)
			VALUES ($1, $2, $3, $4)
		`, "assignment-id-2", account.ID, doctor.ID, now)
		require.NoError(t, err)

		err = repo.ExpireDoctorAssignments(ctx, doctor.ID)
		require.NoError(t, err)

		var validTo time.Time
		err = db.QueryRowContext(ctx, `
			SELECT valid_to FROM doctor_assignments
			WHERE id = $1
		`, "assignment-id-2").Scan(&validTo)
		require.NoError(t, err)
		require.Equal(t, now, validTo)
	})
}

// newTestDoctor - creates a new doctor with a random accountID.
func newTestDoctor(t *testing.T) *domain.Doctor {
	t.Helper()

	account := domain.NewAccount("Dr", "Jones", 45, "dr.jones@example.com")
	require.NoError(t, repo.AddAccount(t.Context(), *account))

	return domain.NewDoctor(account.ID)
}
