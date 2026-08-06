package test

import (
	"context"
	"errors"
	"openhealth/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func TestService_CreateAccount(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully create account", func(t *testing.T) {
		account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")

		deps.tx.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(context.Context) error) error {
				return fn(ctx)
			},
		)

		deps.repository.EXPECT().AddAccount(gomock.Any(), *account).Return(nil)
		// deps.repository.EXPECT().AddOutboxEvent(gomock.Any(), gomock.Any()).Return(nil)

		err := service.CreateAccount(ctx, *account)
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to add account", func(t *testing.T) {
		account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")

		deps.tx.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(context.Context) error) error {
				return fn(ctx)
			},
		)

		deps.repository.EXPECT().AddAccount(gomock.Any(), *account).Return(errors.New("add error"))

		err := service.CreateAccount(ctx, *account)
		require.Error(t, err)
	})

	// t.Run("should handle error when repository fails to add outbox event", func(t *testing.T) {
	// 	account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")

	// 	deps.tx.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(
	// 		func(ctx context.Context, fn func(context.Context) error) error {
	// 			return fn(ctx)
	// 		},
	// 	)

	// 	deps.repository.EXPECT().AddAccount(gomock.Any(), *account).Return(nil)
	// 	// deps.repository.EXPECT().AddOutboxEvent(gomock.Any(), gomock.Any()).Return(errors.New("outbox error"))

	// 	err := service.CreateAccount(ctx, *account)
	// 	require.Error(t, err)
	// })
}

func TestService_GetAccountByID(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully get account", func(t *testing.T) {
		account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")

		deps.repository.EXPECT().GetAccountByID(gomock.Any(), gomock.Any()).Return(account, nil)

		result, err := service.GetAccountByID(ctx, "test-account-id")
		require.NoError(t, err)
		require.Equal(t, account, result)
	})

	t.Run("should handle error when account is not found", func(t *testing.T) {
		deps.repository.EXPECT().GetAccountByID(gomock.Any(), gomock.Any()).Return(nil, nil)

		account, err := service.GetAccountByID(ctx, "non-existent-id")

		require.ErrorIs(t, err, domain.AccountNotFoundError)
		require.Nil(t, account)
	})

	t.Run("should handle error when repository fails to get account", func(t *testing.T) {
		deps.repository.EXPECT().GetAccountByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("repository error"))

		account, err := service.GetAccountByID(ctx, "non-existent-id")

		require.Error(t, err)
		require.Nil(t, account)
	})
}

func TestService_GetAllAccounts(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully return accounts up to the limit", func(t *testing.T) {
		expected := []domain.Account{
			*domain.NewAccount("John", "Doe", 30, "john.doe@example.com"),
			*domain.NewAccount("Jane", "Smith", 25, "jane.smith@example.com"),
		}

		deps.repository.EXPECT().GetAccounts(gomock.Any(), 2).Return(expected, nil)

		accounts, err := service.GetAccounts(ctx, 2)
		require.NoError(t, err)
		require.Len(t, accounts, 2)
		require.Equal(t, expected, accounts)
	})

	t.Run("should return empty slice when no accounts exist", func(t *testing.T) {
		deps.repository.EXPECT().GetAccounts(gomock.Any(), 5).Return([]domain.Account{}, nil)

		accounts, err := service.GetAccounts(ctx, 5)
		require.NoError(t, err)
		require.Empty(t, accounts)
	})

	t.Run("should handle error when repository fails", func(t *testing.T) {
		deps.repository.EXPECT().GetAccounts(gomock.Any(), 10).Return(nil, errors.New("repository error"))

		accounts, err := service.GetAccounts(ctx, 10)
		require.Error(t, err)
		require.Nil(t, accounts)
	})
}

func TestService_UpdateAccount(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully update account", func(t *testing.T) {
		account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")
		account.UpdatedAt = time.Now().UTC()

		deps.repository.EXPECT().UpdateAccount(gomock.Any(), *account).Return(nil)

		err := service.UpdateAccount(ctx, *account)
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to update account", func(t *testing.T) {
		account := domain.NewAccount("John", "Doe", 30, "john.doe@example.com")
		account.UpdatedAt = time.Now().UTC()

		deps.repository.EXPECT().UpdateAccount(gomock.Any(), *account).Return(errors.New("update error"))

		err := service.UpdateAccount(ctx, *account)
		require.Error(t, err)
	})
}

func TestService_ArchiveAccount(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully archive account", func(t *testing.T) {
		deps.repository.EXPECT().ArchiveAccount(gomock.Any(), "account-id").Return(nil)

		err := service.ArchiveAccount(ctx, "account-id")
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to archive account", func(t *testing.T) {
		deps.repository.EXPECT().ArchiveAccount(gomock.Any(), "account-id").Return(errors.New("archive error"))

		err := service.ArchiveAccount(ctx, "account-id")
		require.Error(t, err)
	})
}
