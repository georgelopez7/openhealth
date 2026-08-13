package test

import (
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

func TestRepository_GetDoctorByAccountID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully return doctor by account ID", func(t *testing.T) {
		expected := newTestDoctor(t)

		err := repo.AddDoctor(ctx, *expected)
		require.NoError(t, err)

		doctor, err := repo.GetDoctorByAccountID(ctx, expected.AccountID)
		require.NoError(t, err)
		require.Equal(t, expected.ID, doctor.ID)
		require.Equal(t, expected.AccountID, doctor.AccountID)
		require.Equal(t, domain.DoctorStatusActive, doctor.Status)
	})

	t.Run("should handle case when doctor is not found", func(t *testing.T) {
		doctor, err := repo.GetDoctorByAccountID(ctx, "non-existent-account-id")
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

		err := repo.AddDoctor(ctx, *drJones)
		require.NoError(t, err)

		err = repo.AddDoctor(ctx, *drSmith)
		require.NoError(t, err)

		err = repo.AddDoctor(ctx, *drBrown)
		require.NoError(t, err)

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

		err := repo.AddDoctor(ctx, *expected)
		require.NoError(t, err)

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

func TestRepository_AddDoctorToAccountAssignment(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully add a doctor assignment", func(t *testing.T) {
		doctor := newTestDoctor(t)

		err := repo.AddDoctor(ctx, *doctor)
		require.NoError(t, err)

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
		err = repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		assignment := domain.NewDoctorToAccountAssignment(account.ID, doctor.ID)
		err = repo.AddDoctorToAccountAssignment(ctx, *assignment)
		require.NoError(t, err)

		stored, err := repo.GetDoctorAssignmentByID(ctx, assignment.ID)
		require.NoError(t, err)
		require.Equal(t, assignment.ID, stored.ID)
		require.Equal(t, account.ID, stored.AccountID)
		require.Equal(t, doctor.ID, stored.DoctorID)
		require.Nil(t, stored.ValidTo)
	})
}

func TestRepository_RemoveDoctorAssignment(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully remove a doctor assignment", func(t *testing.T) {
		doctor := newTestDoctor(t)

		err := repo.AddDoctor(ctx, *doctor)
		require.NoError(t, err)

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
		err = repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		assignment := domain.NewDoctorToAccountAssignment(account.ID, doctor.ID)
		err = repo.AddDoctorToAccountAssignment(ctx, *assignment)
		require.NoError(t, err)

		err = repo.RemoveDoctorAssignment(ctx, assignment.ID)
		require.NoError(t, err)

		stored, err := repo.GetDoctorAssignmentByID(ctx, assignment.ID)
		require.NoError(t, err)
		require.NotNil(t, stored.ValidTo)
	})

	t.Run("should handle case when assignment is not found", func(t *testing.T) {
		err := repo.RemoveDoctorAssignment(ctx, "non-existent-id")
		require.NoError(t, err)
	})
}

func TestRepository_RemoveAllDoctorToAccountAssignments(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should remove all open assignments for a doctor", func(t *testing.T) {
		doctor := newTestDoctor(t)

		err := repo.AddDoctor(ctx, *doctor)
		require.NoError(t, err)

		accountOne := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
		accountTwo := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")
		err = repo.AddAccount(ctx, *accountOne)
		require.NoError(t, err)
		err = repo.AddAccount(ctx, *accountTwo)
		require.NoError(t, err)

		assignmentOne := domain.NewDoctorToAccountAssignment(accountOne.ID, doctor.ID)
		assignmentTwo := domain.NewDoctorToAccountAssignment(accountTwo.ID, doctor.ID)
		err = repo.AddDoctorToAccountAssignment(ctx, *assignmentOne)
		require.NoError(t, err)
		err = repo.AddDoctorToAccountAssignment(ctx, *assignmentTwo)
		require.NoError(t, err)

		err = repo.RemoveAllDoctorToAccountAssignments(ctx, doctor.ID)
		require.NoError(t, err)

		assignments, err := repo.GetDoctorAssignmentsByDoctorID(ctx, doctor.ID)
		require.NoError(t, err)
		require.Empty(t, assignments)
	})

	t.Run("should only remove open assignments", func(t *testing.T) {
		repo.ResetAccounts(ctx)

		doctor := newTestDoctor(t)

		err := repo.AddDoctor(ctx, *doctor)
		require.NoError(t, err)

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
		err = repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		now := time.Now().UTC()
		assignment := domain.DoctorAssignment{
			ID:        "assignment-id-1",
			AccountID: account.ID,
			DoctorID:  doctor.ID,
			ValidFrom: now.Add(-time.Hour),
			ValidTo:   &now,
		}
		err = repo.AddDoctorToAccountAssignment(ctx, assignment)
		require.NoError(t, err)

		err = repo.RemoveAllDoctorToAccountAssignments(ctx, doctor.ID)
		require.NoError(t, err)

		stored, err := repo.GetDoctorAssignmentByID(ctx, assignment.ID)
		require.NoError(t, err)
		require.Equal(t, now, *stored.ValidTo)
	})
}

func TestRepository_GetDoctorAssignmentByID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully return assignment", func(t *testing.T) {
		doctor := newTestDoctor(t)

		err := repo.AddDoctor(ctx, *doctor)
		require.NoError(t, err)

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
		err = repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		assignment := domain.NewDoctorToAccountAssignment(account.ID, doctor.ID)
		err = repo.AddDoctorToAccountAssignment(ctx, *assignment)
		require.NoError(t, err)

		stored, err := repo.GetDoctorAssignmentByID(ctx, assignment.ID)
		require.NoError(t, err)
		require.Equal(t, assignment.ID, stored.ID)
		require.Equal(t, account.ID, stored.AccountID)
		require.Equal(t, doctor.ID, stored.DoctorID)
	})

	t.Run("should handle case when assignment is not found", func(t *testing.T) {
		stored, err := repo.GetDoctorAssignmentByID(ctx, "non-existent-id")
		require.NoError(t, err)
		require.Nil(t, stored)
	})
}

func TestRepository_GetDoctorAssignments(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should return assignments up to the limit", func(t *testing.T) {
		doctor := newTestDoctor(t)

		err := repo.AddDoctor(ctx, *doctor)
		require.NoError(t, err)

		accountOne := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
		accountTwo := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")
		err = repo.AddAccount(ctx, *accountOne)
		require.NoError(t, err)
		err = repo.AddAccount(ctx, *accountTwo)
		require.NoError(t, err)

		assignmentOne := domain.NewDoctorToAccountAssignment(accountOne.ID, doctor.ID)
		assignmentTwo := domain.NewDoctorToAccountAssignment(accountTwo.ID, doctor.ID)
		err = repo.AddDoctorToAccountAssignment(ctx, *assignmentOne)
		require.NoError(t, err)
		err = repo.AddDoctorToAccountAssignment(ctx, *assignmentTwo)
		require.NoError(t, err)

		assignments, err := repo.GetDoctorAssignments(ctx, 1)
		require.NoError(t, err)
		require.Len(t, assignments, 1)
	})

	t.Run("should return empty slice when no assignments exist", func(t *testing.T) {
		repo.ResetAccounts(ctx)

		assignments, err := repo.GetDoctorAssignments(ctx, 5)
		require.NoError(t, err)
		require.Empty(t, assignments)
	})
}

func TestRepository_GetDoctorAssignmentsByDoctorID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should return open assignments for a doctor", func(t *testing.T) {
		doctor := newTestDoctor(t)

		err := repo.AddDoctor(ctx, *doctor)
		require.NoError(t, err)

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
		err = repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		assignment := domain.NewDoctorToAccountAssignment(account.ID, doctor.ID)
		err = repo.AddDoctorToAccountAssignment(ctx, *assignment)
		require.NoError(t, err)

		assignments, err := repo.GetDoctorAssignmentsByDoctorID(ctx, doctor.ID)
		require.NoError(t, err)
		require.Len(t, assignments, 1)
		require.Equal(t, assignment.ID, assignments[0].ID)
	})

	t.Run("should not return removed assignments", func(t *testing.T) {
		repo.ResetAccounts(ctx)

		doctor := newTestDoctor(t)

		err := repo.AddDoctor(ctx, *doctor)
		require.NoError(t, err)

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
		err = repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		assignment := domain.NewDoctorToAccountAssignment(account.ID, doctor.ID)
		err = repo.AddDoctorToAccountAssignment(ctx, *assignment)
		require.NoError(t, err)

		err = repo.RemoveDoctorAssignment(ctx, assignment.ID)
		require.NoError(t, err)

		assignments, err := repo.GetDoctorAssignmentsByDoctorID(ctx, doctor.ID)
		require.NoError(t, err)
		require.Empty(t, assignments)
	})
}

// newTestDoctor - creates a new doctor with a random accountID.
func newTestDoctor(t *testing.T) *domain.Doctor {
	t.Helper()

	account := domain.NewAccount("Dr", "Jones", 45, "dr.jones@example.com")
	err := repo.AddAccount(t.Context(), *account)
	require.NoError(t, err)

	return domain.NewDoctor(account.ID)
}
