package test

import (
	"openhealth/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRepository_AddMedicalRecord(t *testing.T) {
	ctx := t.Context()

	repo.ResetMedicalRecords(ctx)
	repo.ResetAccounts(ctx)

	account := newTestAccount(t)
	expected := domain.NewMedicalRecord(account.ID, "Checkup", "Annual physical examination")

	err := repo.AddMedicalRecord(ctx, *expected)
	require.NoError(t, err)

	record, err := repo.GetMedicalRecordByID(ctx, expected.ID)
	require.NoError(t, err)
	require.Equal(t, expected.ID, record.ID)
	require.Equal(t, expected.AccountID, record.AccountID)
	require.Equal(t, expected.Title, record.Title)
	require.Equal(t, expected.Description, record.Description)
	require.Equal(t, domain.MedicalRecordStatusActive, record.Status)
}

func TestRepository_GetMedicalRecordByID(t *testing.T) {
	ctx := t.Context()

	repo.ResetMedicalRecords(ctx)
	repo.ResetAccounts(ctx)

	t.Run("should successfully return medical record", func(t *testing.T) {
		account := newTestAccount(t)
		expected := domain.NewMedicalRecord(account.ID, "Lab Results", "Blood work results")

		err := repo.AddMedicalRecord(ctx, *expected)
		require.NoError(t, err)

		record, err := repo.GetMedicalRecordByID(ctx, expected.ID)
		require.NoError(t, err)
		require.Equal(t, expected.ID, record.ID)
		require.Equal(t, expected.AccountID, record.AccountID)
		require.Equal(t, expected.Title, record.Title)
		require.Equal(t, expected.Description, record.Description)
		require.Equal(t, domain.MedicalRecordStatusActive, record.Status)
	})

	t.Run("should handle case when medical record is not found", func(t *testing.T) {
		record, err := repo.GetMedicalRecordByID(ctx, "non-existent-id")
		require.NoError(t, err)
		require.Nil(t, record)
	})
}

func TestRepository_GetMedicalRecords(t *testing.T) {
	ctx := t.Context()

	repo.ResetMedicalRecords(ctx)
	repo.ResetAccounts(ctx)

	t.Run("should return all medical records", func(t *testing.T) {
		account := newTestAccount(t)
		expected := domain.NewMedicalRecord(account.ID, "Lab Results", "Blood work results")

		err := repo.AddMedicalRecord(ctx, *expected)
		require.NoError(t, err)

		records, err := repo.GetMedicalRecords(ctx)
		require.NoError(t, err)
		require.Len(t, records, 1)
		require.Equal(t, expected.ID, records[0].ID)
		require.Equal(t, expected.AccountID, records[0].AccountID)
		require.Equal(t, expected.Title, records[0].Title)
	})
}

func TestRepository_UpdateMedicalRecord(t *testing.T) {
	ctx := t.Context()

	repo.ResetMedicalRecords(ctx)
	repo.ResetAccounts(ctx)

	t.Run("should successfully update medical record", func(t *testing.T) {
		account := newTestAccount(t)
		record := domain.NewMedicalRecord(account.ID, "Checkup", "Annual physical examination")

		err := repo.AddMedicalRecord(ctx, *record)
		require.NoError(t, err)

		record.Title = "Updated Checkup"
		record.Description = "Updated annual physical examination"
		record.UpdatedAt = time.Now().UTC()

		err = repo.UpdateMedicalRecord(ctx, *record)
		require.NoError(t, err)

		updated, err := repo.GetMedicalRecordByID(ctx, record.ID)
		require.NoError(t, err)
		require.Equal(t, "Updated Checkup", updated.Title)
		require.Equal(t, "Updated annual physical examination", updated.Description)
		require.True(t, updated.UpdatedAt.After(updated.CreatedAt) || updated.UpdatedAt.Equal(updated.CreatedAt))
	})

	t.Run("should handle case when medical record is not found", func(t *testing.T) {
		record := domain.NewMedicalRecord("non-existent-account", "Ghost", "Ghost record")
		record.ID = "non-existent-id"

		err := repo.UpdateMedicalRecord(ctx, *record)
		require.NoError(t, err)
	})
}

func TestRepository_ArchiveMedicalRecord(t *testing.T) {
	ctx := t.Context()

	repo.ResetMedicalRecords(ctx)
	repo.ResetAccounts(ctx)

	t.Run("should successfully archive medical record", func(t *testing.T) {
		account := newTestAccount(t)
		record := domain.NewMedicalRecord(account.ID, "Checkup", "Annual physical examination")

		err := repo.AddMedicalRecord(ctx, *record)
		require.NoError(t, err)

		err = repo.ArchiveMedicalRecord(ctx, record.ID)
		require.NoError(t, err)

		archived, err := repo.GetMedicalRecordByID(ctx, record.ID)
		require.NoError(t, err)
		require.Equal(t, domain.MedicalRecordStatusArchived, archived.Status)
	})

	t.Run("should handle case when medical record is not found", func(t *testing.T) {
		err := repo.ArchiveMedicalRecord(ctx, "non-existent-id")
		require.NoError(t, err)
	})
}

// newTestAccount - creates a new account for use in repository tests.
func newTestAccount(t *testing.T) *domain.Account {
	t.Helper()

	account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com", "")
	err := repo.AddAccount(t.Context(), *account)
	require.NoError(t, err)

	return account
}
