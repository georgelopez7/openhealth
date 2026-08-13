package test

import (
	"openhealth/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRepository_AddNurse(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	expected := newTestNurse(t)

	err := repo.AddNurse(ctx, *expected)
	require.NoError(t, err)

	nurse, err := repo.GetNurseByID(ctx, expected.ID)
	require.NoError(t, err)
	require.Equal(t, expected.ID, nurse.ID)
	require.Equal(t, expected.AccountID, nurse.AccountID)
	require.Equal(t, domain.NurseStatusActive, nurse.Status)
}

func TestRepository_GetNurseByID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully return nurse", func(t *testing.T) {
		expected := newTestNurse(t)

		err := repo.AddNurse(ctx, *expected)
		require.NoError(t, err)

		nurse, err := repo.GetNurseByID(ctx, expected.ID)
		require.NoError(t, err)
		require.Equal(t, expected.ID, nurse.ID)
		require.Equal(t, expected.AccountID, nurse.AccountID)
		require.Equal(t, domain.NurseStatusActive, nurse.Status)
	})

	t.Run("should handle case when nurse is not found", func(t *testing.T) {
		nurse, err := repo.GetNurseByID(ctx, "non-existent-id")
		require.NoError(t, err)
		require.Nil(t, nurse)
	})
}

func TestRepository_GetNurseByAccountID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully return nurse by account ID", func(t *testing.T) {
		expected := newTestNurse(t)

		err := repo.AddNurse(ctx, *expected)
		require.NoError(t, err)

		nurse, err := repo.GetNurseByAccountID(ctx, expected.AccountID)
		require.NoError(t, err)
		require.Equal(t, expected.ID, nurse.ID)
		require.Equal(t, expected.AccountID, nurse.AccountID)
		require.Equal(t, domain.NurseStatusActive, nurse.Status)
	})

	t.Run("should handle case when nurse is not found", func(t *testing.T) {
		nurse, err := repo.GetNurseByAccountID(ctx, "non-existent-account-id")
		require.NoError(t, err)
		require.Nil(t, nurse)
	})
}

func TestRepository_GetNurses(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should return nurses up to the limit", func(t *testing.T) {
		jones := newTestNurse(t)
		smith := newTestNurse(t)
		brown := newTestNurse(t)

		err := repo.AddNurse(ctx, *jones)
		require.NoError(t, err)

		err = repo.AddNurse(ctx, *smith)
		require.NoError(t, err)

		err = repo.AddNurse(ctx, *brown)
		require.NoError(t, err)

		nurses, err := repo.GetNurses(ctx, 2)
		require.NoError(t, err)
		require.Len(t, nurses, 2)

		createdIDs := []string{jones.ID, smith.ID, brown.ID}

		require.Contains(t, createdIDs, nurses[0].ID)
		require.Contains(t, createdIDs, nurses[1].ID)
	})

	t.Run("should return all nurses when limit is greater than count", func(t *testing.T) {
		repo.ResetAccounts(ctx)

		expected := newTestNurse(t)

		err := repo.AddNurse(ctx, *expected)
		require.NoError(t, err)

		nurses, err := repo.GetNurses(ctx, 10)
		require.NoError(t, err)
		require.Len(t, nurses, 1)
		require.Equal(t, expected.ID, nurses[0].ID)
	})

	t.Run("should return empty slice when no nurses exist", func(t *testing.T) {
		repo.ResetAccounts(ctx)

		nurses, err := repo.GetNurses(ctx, 5)
		require.NoError(t, err)
		require.Empty(t, nurses)
	})
}

func TestRepository_UpdateNurseStatusByID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully update nurse status", func(t *testing.T) {
		nurse := newTestNurse(t)

		err := repo.AddNurse(ctx, *nurse)
		require.NoError(t, err)

		err = repo.UpdateNurseStatusByID(ctx, nurse.ID, domain.NurseStatusArchived)
		require.NoError(t, err)

		archived, err := repo.GetNurseByID(ctx, nurse.ID)
		require.NoError(t, err)
		require.Equal(t, domain.NurseStatusArchived, archived.Status)
	})

	t.Run("should handle case when nurse is not found", func(t *testing.T) {
		err := repo.UpdateNurseStatusByID(ctx, "non-existent-id", domain.NurseStatusArchived)
		require.NoError(t, err)
	})
}

func TestRepository_AddNurseToHospitalAssignment(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)
	repo.ResetHospitals(ctx)

	t.Run("should successfully add a nurse to hospital assignment", func(t *testing.T) {
		nurse := newTestNurse(t)
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddNurse(ctx, *nurse)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		assignment := domain.NewNurseToHospitalAssignment(nurse.ID, hospital.ID)
		err = repo.AddNurseToHospitalAssignment(ctx, *assignment)
		require.NoError(t, err)

		stored, err := repo.GetNurseToHospitalAssignmentByID(ctx, assignment.ID)
		require.NoError(t, err)
		require.Equal(t, assignment.ID, stored.ID)
		require.Equal(t, nurse.ID, stored.NurseID)
		require.Equal(t, hospital.ID, stored.HospitalID)
		require.Nil(t, stored.ValidTo)
	})
}

func TestRepository_RemoveNurseToHospitalAssignment(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)
	repo.ResetHospitals(ctx)

	t.Run("should successfully remove a nurse to hospital assignment", func(t *testing.T) {
		nurse := newTestNurse(t)
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddNurse(ctx, *nurse)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		assignment := domain.NewNurseToHospitalAssignment(nurse.ID, hospital.ID)
		err = repo.AddNurseToHospitalAssignment(ctx, *assignment)
		require.NoError(t, err)

		err = repo.RemoveNurseToHospitalAssignment(ctx, assignment.ID)
		require.NoError(t, err)

		stored, err := repo.GetNurseToHospitalAssignmentByID(ctx, assignment.ID)
		require.NoError(t, err)
		require.NotNil(t, stored.ValidTo)
	})

	t.Run("should handle case when assignment is not found", func(t *testing.T) {
		err := repo.RemoveNurseToHospitalAssignment(ctx, "non-existent-id")
		require.NoError(t, err)
	})
}

func TestRepository_RemoveAllNurseToHospitalAssignments(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)
	repo.ResetHospitals(ctx)

	t.Run("should remove all open assignments for a nurse", func(t *testing.T) {
		nurse := newTestNurse(t)
		hospitalOne := domain.NewHospital("St. Test Hospital")
		hospitalTwo := domain.NewHospital("General Hospital")

		err := repo.AddNurse(ctx, *nurse)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *hospitalOne)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *hospitalTwo)
		require.NoError(t, err)

		assignmentOne := domain.NewNurseToHospitalAssignment(nurse.ID, hospitalOne.ID)
		assignmentTwo := domain.NewNurseToHospitalAssignment(nurse.ID, hospitalTwo.ID)
		err = repo.AddNurseToHospitalAssignment(ctx, *assignmentOne)
		require.NoError(t, err)
		err = repo.AddNurseToHospitalAssignment(ctx, *assignmentTwo)
		require.NoError(t, err)

		err = repo.RemoveAllNurseToHospitalAssignments(ctx, nurse.ID)
		require.NoError(t, err)

		assignments, err := repo.GetNurseToHospitalAssignmentsByNurseID(ctx, nurse.ID)
		require.NoError(t, err)
		require.Empty(t, assignments)
	})

	t.Run("should only remove open assignments", func(t *testing.T) {
		repo.ResetAccounts(ctx)
		repo.ResetHospitals(ctx)

		nurse := newTestNurse(t)
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddNurse(ctx, *nurse)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		now := time.Now().UTC()
		assignment := domain.NurseToHospitalAssignment{
			ID:         "assignment-id-1",
			NurseID:    nurse.ID,
			HospitalID: hospital.ID,
			ValidFrom:  now.Add(-time.Hour),
			ValidTo:    &now,
		}
		err = repo.AddNurseToHospitalAssignment(ctx, assignment)
		require.NoError(t, err)

		err = repo.RemoveAllNurseToHospitalAssignments(ctx, nurse.ID)
		require.NoError(t, err)

		stored, err := repo.GetNurseToHospitalAssignmentByID(ctx, assignment.ID)
		require.NoError(t, err)
		require.Equal(t, now, *stored.ValidTo)
	})
}

func TestRepository_GetNurseToHospitalAssignmentByID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)
	repo.ResetHospitals(ctx)

	t.Run("should successfully return assignment", func(t *testing.T) {
		nurse := newTestNurse(t)
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddNurse(ctx, *nurse)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		assignment := domain.NewNurseToHospitalAssignment(nurse.ID, hospital.ID)
		err = repo.AddNurseToHospitalAssignment(ctx, *assignment)
		require.NoError(t, err)

		stored, err := repo.GetNurseToHospitalAssignmentByID(ctx, assignment.ID)
		require.NoError(t, err)
		require.Equal(t, assignment.ID, stored.ID)
		require.Equal(t, nurse.ID, stored.NurseID)
		require.Equal(t, hospital.ID, stored.HospitalID)
	})

	t.Run("should handle case when assignment is not found", func(t *testing.T) {
		stored, err := repo.GetNurseToHospitalAssignmentByID(ctx, "non-existent-id")
		require.NoError(t, err)
		require.Nil(t, stored)
	})
}

func TestRepository_GetNurseToHospitalAssignments(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)
	repo.ResetHospitals(ctx)

	t.Run("should return assignments up to the limit", func(t *testing.T) {
		nurse := newTestNurse(t)
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddNurse(ctx, *nurse)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		assignmentOne := domain.NewNurseToHospitalAssignment(nurse.ID, hospital.ID)
		assignmentTwo := domain.NewNurseToHospitalAssignment(nurse.ID, hospital.ID)
		err = repo.AddNurseToHospitalAssignment(ctx, *assignmentOne)
		require.NoError(t, err)
		err = repo.AddNurseToHospitalAssignment(ctx, *assignmentTwo)
		require.NoError(t, err)

		assignments, err := repo.GetNurseToHospitalAssignments(ctx, 1)
		require.NoError(t, err)
		require.Len(t, assignments, 1)
	})

	t.Run("should return empty slice when no assignments exist", func(t *testing.T) {
		repo.ResetAccounts(ctx)
		repo.ResetHospitals(ctx)

		assignments, err := repo.GetNurseToHospitalAssignments(ctx, 5)
		require.NoError(t, err)
		require.Empty(t, assignments)
	})
}

func TestRepository_GetNurseToHospitalAssignmentsByNurseID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)
	repo.ResetHospitals(ctx)

	t.Run("should return open assignments for a nurse", func(t *testing.T) {
		nurse := newTestNurse(t)
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddNurse(ctx, *nurse)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		assignment := domain.NewNurseToHospitalAssignment(nurse.ID, hospital.ID)
		err = repo.AddNurseToHospitalAssignment(ctx, *assignment)
		require.NoError(t, err)

		assignments, err := repo.GetNurseToHospitalAssignmentsByNurseID(ctx, nurse.ID)
		require.NoError(t, err)
		require.Len(t, assignments, 1)
		require.Equal(t, assignment.ID, assignments[0].ID)
	})

	t.Run("should not return removed assignments", func(t *testing.T) {
		repo.ResetAccounts(ctx)
		repo.ResetHospitals(ctx)

		nurse := newTestNurse(t)
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddNurse(ctx, *nurse)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		assignment := domain.NewNurseToHospitalAssignment(nurse.ID, hospital.ID)
		err = repo.AddNurseToHospitalAssignment(ctx, *assignment)
		require.NoError(t, err)

		err = repo.RemoveNurseToHospitalAssignment(ctx, assignment.ID)
		require.NoError(t, err)

		assignments, err := repo.GetNurseToHospitalAssignmentsByNurseID(ctx, nurse.ID)
		require.NoError(t, err)
		require.Empty(t, assignments)
	})
}

func TestRepository_GetHospitalIDByNurseID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)
	repo.ResetHospitals(ctx)

	t.Run("should return the currently assigned hospital ID", func(t *testing.T) {
		nurse := newTestNurse(t)
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddNurse(ctx, *nurse)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		assignment := domain.NewNurseToHospitalAssignment(nurse.ID, hospital.ID)
		err = repo.AddNurseToHospitalAssignment(ctx, *assignment)
		require.NoError(t, err)

		hospitalID, err := repo.GetHospitalIDByNurseID(ctx, nurse.ID)
		require.NoError(t, err)
		require.Equal(t, hospital.ID, hospitalID)
	})

	t.Run("should return empty when no hospital is assigned", func(t *testing.T) {
		nurse := newTestNurse(t)

		err := repo.AddNurse(ctx, *nurse)
		require.NoError(t, err)

		hospitalID, err := repo.GetHospitalIDByNurseID(ctx, nurse.ID)
		require.NoError(t, err)
		require.Empty(t, hospitalID)
	})
}

func TestRepository_GetNursesByHospitalID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)
	repo.ResetHospitals(ctx)

	t.Run("should return open assignments for a hospital", func(t *testing.T) {
		nurse := newTestNurse(t)
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddNurse(ctx, *nurse)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		assignment := domain.NewNurseToHospitalAssignment(nurse.ID, hospital.ID)
		err = repo.AddNurseToHospitalAssignment(ctx, *assignment)
		require.NoError(t, err)

		assignments, err := repo.GetNursesByHospitalID(ctx, hospital.ID)
		require.NoError(t, err)
		require.Len(t, assignments, 1)
		require.Equal(t, assignment.ID, assignments[0].ID)
	})

	t.Run("should not return removed assignments", func(t *testing.T) {
		repo.ResetAccounts(ctx)
		repo.ResetHospitals(ctx)

		nurse := newTestNurse(t)
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddNurse(ctx, *nurse)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		assignment := domain.NewNurseToHospitalAssignment(nurse.ID, hospital.ID)
		err = repo.AddNurseToHospitalAssignment(ctx, *assignment)
		require.NoError(t, err)

		err = repo.RemoveNurseToHospitalAssignment(ctx, assignment.ID)
		require.NoError(t, err)

		assignments, err := repo.GetNursesByHospitalID(ctx, hospital.ID)
		require.NoError(t, err)
		require.Empty(t, assignments)
	})
}

// newTestNurse - creates a new nurse with a random accountID.
func newTestNurse(t *testing.T) *domain.Nurse {
	t.Helper()

	account := domain.NewAccount("Nurse", "Joy", 30, "nurse.joy@example.com")
	err := repo.AddAccount(t.Context(), *account)
	require.NoError(t, err)

	return domain.NewNurse(account.ID)
}
