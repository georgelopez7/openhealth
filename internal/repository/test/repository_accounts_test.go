package test

import (
	"openhealth/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRepository_AddAccount(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	expected := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")

	err := repo.AddAccount(ctx, *expected)
	require.NoError(t, err)

	account, err := repo.GetAccountByID(ctx, expected.ID)
	require.NoError(t, err)
	require.Equal(t, expected.ID, account.ID)
	require.Equal(t, expected.FirstName, account.FirstName)
	require.Equal(t, expected.LastName, account.LastName)
	require.Equal(t, expected.Age, account.Age)
	require.Equal(t, expected.Email, account.Email)
	require.Equal(t, domain.AccountStatusActive, account.Status)
}

func TestRepository_GetAccountByID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully return account", func(t *testing.T) {
		expected := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")

		err := repo.AddAccount(ctx, *expected)
		require.NoError(t, err)

		account, err := repo.GetAccountByID(ctx, expected.ID)
		require.NoError(t, err)
		require.Equal(t, expected.ID, account.ID)
		require.Equal(t, expected.FirstName, account.FirstName)
		require.Equal(t, expected.LastName, account.LastName)
		require.Equal(t, expected.Age, account.Age)
		require.Equal(t, expected.Email, account.Email)
		require.Equal(t, domain.AccountStatusActive, account.Status)
	})

	t.Run("should handle case when account is not found", func(t *testing.T) {
		account, err := repo.GetAccountByID(ctx, "non-existent-id")
		require.NoError(t, err)
		require.Nil(t, account)
	})
}

func TestRepository_GetAccounts(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should return accounts up to the limit", func(t *testing.T) {
		john := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")
		jane := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
		bob := domain.NewAccount("Bob", "Brown", 40, "bob.brown@example.com")

		err := repo.AddAccount(ctx, *john)
		require.NoError(t, err)
		err = repo.AddAccount(ctx, *jane)
		require.NoError(t, err)
		err = repo.AddAccount(ctx, *bob)
		require.NoError(t, err)

		accounts, err := repo.GetAccounts(ctx, 2)
		require.NoError(t, err)
		require.Len(t, accounts, 2)

		createdIDs := []string{john.ID, jane.ID, bob.ID}

		require.Contains(t, createdIDs, accounts[0].ID)
		require.Contains(t, createdIDs, accounts[1].ID)
	})

	t.Run("should return all accounts when limit is greater than count", func(t *testing.T) {
		repo.ResetAccounts(ctx)

		expected := domain.NewAccount("Alice", "Wonder", 28, "alice@example.com")
		err := repo.AddAccount(ctx, *expected)
		require.NoError(t, err)

		accounts, err := repo.GetAccounts(ctx, 10)
		require.NoError(t, err)
		require.Len(t, accounts, 1)
		require.Equal(t, expected.ID, accounts[0].ID)
	})

	t.Run("should return empty slice when no accounts exist", func(t *testing.T) {
		repo.ResetAccounts(ctx)

		accounts, err := repo.GetAccounts(ctx, 5)
		require.NoError(t, err)
		require.Empty(t, accounts)
	})
}

func TestRepository_UpdateAccount(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully update account", func(t *testing.T) {
		account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")

		err := repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		account.FirstName = "Jonathan"
		account.LastName = "Dorian"
		account.Age = 31
		account.Email = "jonathan.dorian@example.com"
		account.UpdatedAt = time.Now().UTC()

		err = repo.UpdateAccount(ctx, *account)
		require.NoError(t, err)

		updated, err := repo.GetAccountByID(ctx, account.ID)
		require.NoError(t, err)
		require.Equal(t, "Jonathan", updated.FirstName)
		require.Equal(t, "Dorian", updated.LastName)
		require.Equal(t, 31, updated.Age)
		require.Equal(t, "jonathan.dorian@example.com", updated.Email)
		require.True(t, updated.UpdatedAt.After(updated.CreatedAt) || updated.UpdatedAt.Equal(updated.CreatedAt))
	})

	t.Run("should handle case when account is not found", func(t *testing.T) {
		account := domain.NewAccount("Ghost", "User", 99, "ghost@example.com")
		account.ID = "non-existent-id"

		err := repo.UpdateAccount(ctx, *account)
		require.NoError(t, err)
	})
}

func TestRepository_GetDoctorIDByAccountID(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should return assigned doctor id", func(t *testing.T) {
		doctor := newTestDoctor(t)
		err := repo.AddDoctor(ctx, *doctor)
		require.NoError(t, err)

		account := domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com")
		err = repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		assignment := domain.NewDoctorAssignment(account.ID, doctor.ID)
		err = repo.AddDoctorToAccountAssignment(ctx, *assignment)
		require.NoError(t, err)

		doctorID, err := repo.GetDoctorIDByAccountID(ctx, account.ID)
		require.NoError(t, err)
		require.Equal(t, doctor.ID, doctorID)
	})

	t.Run("should return empty string when no doctor is assigned", func(t *testing.T) {
		account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")
		err := repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		doctorID, err := repo.GetDoctorIDByAccountID(ctx, account.ID)
		require.NoError(t, err)
		require.Empty(t, doctorID)
	})

	t.Run("should return empty string when assignment is expired", func(t *testing.T) {
		doctor := newTestDoctor(t)

		err := repo.AddDoctor(ctx, *doctor)
		require.NoError(t, err)

		account := domain.NewAccount("Bob", "Brown", 40, "bob.brown@example.com")

		err = repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		assignment := domain.NewDoctorAssignment(account.ID, doctor.ID)
		err = repo.AddDoctorToAccountAssignment(ctx, *assignment)
		require.NoError(t, err)

		err = repo.RemoveDoctorAssignment(ctx, assignment.ID)
		require.NoError(t, err)

		doctorID, err := repo.GetDoctorIDByAccountID(ctx, account.ID)
		require.NoError(t, err)
		require.Empty(t, doctorID)
	})
}

func TestRepository_ArchiveAccount(t *testing.T) {
	ctx := t.Context()

	repo.ResetAccounts(ctx)

	t.Run("should successfully archive account", func(t *testing.T) {
		account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")

		err := repo.AddAccount(ctx, *account)
		require.NoError(t, err)

		err = repo.ArchiveAccount(ctx, account.ID)
		require.NoError(t, err)

		archived, err := repo.GetAccountByID(ctx, account.ID)
		require.NoError(t, err)
		require.Equal(t, domain.AccountStatusArchived, archived.Status)
	})

	t.Run("should handle case when account is not found", func(t *testing.T) {
		err := repo.ArchiveAccount(ctx, "non-existent-id")
		require.NoError(t, err)
	})
}
