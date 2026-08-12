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

	repo.ResetHospitals(ctx)

	t.Run("should return hospitals up to the limit", func(t *testing.T) {
		general := domain.NewHospital("St. Test Hospital")
		regional := domain.NewHospital("St. Test Hospital")
		childrens := domain.NewHospital("St. Test Hospital")

		require.NoError(t, repo.AddHospital(ctx, *general))
		require.NoError(t, repo.AddHospital(ctx, *regional))
		require.NoError(t, repo.AddHospital(ctx, *childrens))

		hospitals, err := repo.GetHospitals(ctx, 2)
		require.NoError(t, err)
		require.Len(t, hospitals, 2)

		createdIDs := []string{general.ID, regional.ID, childrens.ID}

		require.Contains(t, createdIDs, hospitals[0].ID)
		require.Contains(t, createdIDs, hospitals[1].ID)
	})

	t.Run("should return all hospitals when limit is greater than count", func(t *testing.T) {
		repo.ResetHospitals(ctx)

		expected := domain.NewHospital("St. Test Hospital")
		require.NoError(t, repo.AddHospital(ctx, *expected))

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

