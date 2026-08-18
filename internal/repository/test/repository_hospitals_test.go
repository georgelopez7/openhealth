package test

import (
	"openhealth/internal/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepository_AddHospital(t *testing.T) {
	ctx := t.Context()

	repo.ResetHospitals(ctx)

	expected := domain.NewHospital("St. Test Hospital")

	err := repo.AddHospital(ctx, *expected)
	require.NoError(t, err)

	hospital, err := repo.GetHospitalByID(ctx, expected.ID)
	require.NoError(t, err)
	require.Equal(t, expected.ID, hospital.ID)
	require.Equal(t, expected.Name, hospital.Name)
	require.Equal(t, domain.HospitalStatusActive, hospital.Status)
}

func TestRepository_GetHospitalByID(t *testing.T) {
	ctx := t.Context()

	repo.ResetHospitals(ctx)

	t.Run("should successfully return hospital", func(t *testing.T) {
		expected := domain.NewHospital("St. Test Hospital")

		err := repo.AddHospital(ctx, *expected)
		require.NoError(t, err)

		hospital, err := repo.GetHospitalByID(ctx, expected.ID)
		require.NoError(t, err)
		require.Equal(t, expected.ID, hospital.ID)
		require.Equal(t, expected.Name, hospital.Name)
		require.Equal(t, domain.HospitalStatusActive, hospital.Status)
	})

	t.Run("should handle case when hospital is not found", func(t *testing.T) {
		hospital, err := repo.GetHospitalByID(ctx, "non-existent-id")
		require.NoError(t, err)
		require.Nil(t, hospital)
	})
}

func TestRepository_GetHospitals(t *testing.T) {
	ctx := t.Context()

	t.Run("should return hospitals up to the limit", func(t *testing.T) {
		repo.ResetHospitals(ctx)

		general := domain.NewHospital("General Hospital")
		regional := domain.NewHospital("Regional Hospital")
		children := domain.NewHospital("Children's Hospital")

		err := repo.AddHospital(ctx, *general)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *regional)
		require.NoError(t, err)

		err = repo.AddHospital(ctx, *children)
		require.NoError(t, err)

		hospitals, err := repo.GetHospitals(ctx, 2)
		require.NoError(t, err)
		require.Len(t, hospitals, 2)

		createdIDs := []string{general.ID, regional.ID, children.ID}

		require.Contains(t, createdIDs, hospitals[0].ID)
		require.Contains(t, createdIDs, hospitals[1].ID)
	})

	t.Run("should return all hospitals when limit is greater than count", func(t *testing.T) {
		repo.ResetHospitals(ctx)

		expected := domain.NewHospital("St. Test Hospital")
		err := repo.AddHospital(ctx, *expected)
		require.NoError(t, err)

		hospitals, err := repo.GetHospitals(ctx, 10)
		require.NoError(t, err)
		require.Len(t, hospitals, 1)
		require.Equal(t, expected.ID, hospitals[0].ID)
	})

	t.Run("should return empty slice when no hospitals exist", func(t *testing.T) {
		repo.ResetHospitals(ctx)

		hospitals, err := repo.GetHospitals(ctx, 5)
		require.NoError(t, err)
		require.Empty(t, hospitals)
	})
}

func TestRepository_UpdateHospital(t *testing.T) {
	ctx := t.Context()

	repo.ResetHospitals(ctx)

	t.Run("should successfully update hospital", func(t *testing.T) {
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		hospital.Name = "Updated Hospital Name"
		hospital.Status = domain.HospitalStatusArchived

		err = repo.UpdateHospital(ctx, *hospital)
		require.NoError(t, err)

		updated, err := repo.GetHospitalByID(ctx, hospital.ID)
		require.NoError(t, err)
		require.Equal(t, "Updated Hospital Name", updated.Name)
		require.Equal(t, domain.HospitalStatusArchived, updated.Status)
	})

	t.Run("should handle case when hospital is not found", func(t *testing.T) {
		hospital := domain.NewHospital("Ghost Hospital")

		err := repo.UpdateHospital(ctx, *hospital)
		require.NoError(t, err)
	})
}

func TestRepository_ArchiveHospital(t *testing.T) {
	ctx := t.Context()

	repo.ResetHospitals(ctx)

	t.Run("should successfully archive hospital", func(t *testing.T) {
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		err = repo.ArchiveHospital(ctx, hospital.ID)
		require.NoError(t, err)

		archived, err := repo.GetHospitalByID(ctx, hospital.ID)
		require.NoError(t, err)
		require.Equal(t, domain.HospitalStatusArchived, archived.Status)
	})

	t.Run("should handle case when hospital is not found", func(t *testing.T) {
		err := repo.ArchiveHospital(ctx, "non-existent-id")
		require.NoError(t, err)
	})
}

func TestRepository_AddHospitalAssignment(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully add a hospital assignment", func(t *testing.T) {
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com", "")

		err = repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		assignment := domain.NewHospitalAssignment(account.ID, hospital.ID)

		err = repo.AddHospitalAssignment(ctx, *assignment)
		require.NoError(t, err)

		stored, err := repo.GetHospitalAssignmentByID(ctx, assignment.ID)
		require.NoError(t, err)
		require.Equal(t, assignment.ID, stored.ID)
		require.Equal(t, account.ID, stored.AccountID)
		require.Equal(t, hospital.ID, stored.HospitalID)
		require.Nil(t, stored.ValidTo)
	})
}

func TestRepository_RemoveHospitalToAccountAssignment(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully remove a hospital to account assignment", func(t *testing.T) {
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com", "")

		err = repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		assignment := domain.NewHospitalAssignment(account.ID, hospital.ID)

		err = repo.AddHospitalAssignment(ctx, *assignment)
		require.NoError(t, err)

		err = repo.RemoveHospitalToAccountAssignment(ctx, assignment.ID)
		require.NoError(t, err)

		stored, err := repo.GetHospitalAssignmentByID(ctx, assignment.ID)
		require.NoError(t, err)
		require.NotNil(t, stored.ValidTo)
	})

	t.Run("should handle case when assignment is not found", func(t *testing.T) {
		err := repo.RemoveHospitalToAccountAssignment(ctx, "non-existent-id")
		require.NoError(t, err)
	})
}

func TestRepository_GetHospitalAssignmentByID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully return assignment", func(t *testing.T) {
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com", "")

		err = repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		assignment := domain.NewHospitalAssignment(account.ID, hospital.ID)

		err = repo.AddHospitalAssignment(ctx, *assignment)
		require.NoError(t, err)

		stored, err := repo.GetHospitalAssignmentByID(ctx, assignment.ID)
		require.NoError(t, err)
		require.Equal(t, assignment.ID, stored.ID)
		require.Equal(t, account.ID, stored.AccountID)
		require.Equal(t, hospital.ID, stored.HospitalID)
	})

	t.Run("should handle case when assignment is not found", func(t *testing.T) {
		stored, err := repo.GetHospitalAssignmentByID(ctx, "non-existent-id")
		require.NoError(t, err)
		require.Nil(t, stored)
	})
}

func TestRepository_GetHospitalAssignments(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should return assignments up to the limit", func(t *testing.T) {
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		account1 := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com", "")
		account2 := domain.NewAccount("John", "Doe", 30, "john.doe@example.com", "")

		err = repo.AddAccount(ctx, *account1)
		require.NoError(t, err)

		err = repo.AddAccount(ctx, *account2)
		require.NoError(t, err)

		assignmentOne := domain.NewHospitalAssignment(account1.ID, hospital.ID)
		assignmentTwo := domain.NewHospitalAssignment(account2.ID, hospital.ID)

		err = repo.AddHospitalAssignment(ctx, *assignmentOne)
		require.NoError(t, err)

		err = repo.AddHospitalAssignment(ctx, *assignmentTwo)
		require.NoError(t, err)

		assignments, err := repo.GetHospitalAssignments(ctx, 1)
		require.NoError(t, err)
		require.Len(t, assignments, 1)
	})

	t.Run("should return empty slice when no assignments exist", func(t *testing.T) {
		repo.ResetHospitals(ctx)

		assignments, err := repo.GetHospitalAssignments(ctx, 5)
		require.NoError(t, err)
		require.Empty(t, assignments)
	})
}

func TestRepository_GetHospitalIDByAccountID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should return the currently assigned hospital ID", func(t *testing.T) {
		hospital := domain.NewHospital("St. Test Hospital")

		err := repo.AddHospital(ctx, *hospital)
		require.NoError(t, err)

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com", "")

		err = repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		assignment := domain.NewHospitalAssignment(account.ID, hospital.ID)

		err = repo.AddHospitalAssignment(ctx, *assignment)
		require.NoError(t, err)

		hospitalID, err := repo.GetHospitalIDByAccountID(ctx, account.ID)

		require.NoError(t, err)
		require.Equal(t, hospital.ID, hospitalID)
	})

	t.Run("should return empty when no hospital is assigned", func(t *testing.T) {
		account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com", "")

		err := repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		hospitalID, err := repo.GetHospitalIDByAccountID(ctx, account.ID)
		require.NoError(t, err)
		require.Empty(t, hospitalID)
	})
}
