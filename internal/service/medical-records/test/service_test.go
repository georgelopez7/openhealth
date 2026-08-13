package test

import (
	"errors"
	"openhealth/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func TestService_CreateMedicalRecord(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully create medical record", func(t *testing.T) {
		record := domain.NewMedicalRecord("account-id-1", "Checkup", "Annual physical examination")

		deps.repository.EXPECT().AddMedicalRecord(gomock.Any(), *record).Return(nil)

		err := service.CreateMedicalRecord(ctx, *record)
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to add medical record", func(t *testing.T) {
		record := domain.NewMedicalRecord("account-id-1", "Checkup", "Annual physical examination")

		deps.repository.EXPECT().AddMedicalRecord(gomock.Any(), *record).Return(errors.New("add error"))

		err := service.CreateMedicalRecord(ctx, *record)
		require.Error(t, err)
	})
}

func TestService_GetMedicalRecordByID(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully get medical record", func(t *testing.T) {
		record := domain.NewMedicalRecord("account-id-1", "Checkup", "Annual physical examination")

		deps.repository.EXPECT().GetMedicalRecordByID(gomock.Any(), "record-id-1").Return(record, nil)

		result, err := service.GetMedicalRecordByID(ctx, "record-id-1")
		require.NoError(t, err)
		require.Equal(t, record, result)
	})

	t.Run("should handle error when medical record is not found", func(t *testing.T) {
		deps.repository.EXPECT().GetMedicalRecordByID(gomock.Any(), "record-id-1").Return(nil, nil)

		record, err := service.GetMedicalRecordByID(ctx, "record-id-1")
		require.ErrorIs(t, err, domain.MedicalRecordNotFoundError)
		require.Nil(t, record)
	})

	t.Run("should handle error when repository fails to get medical record", func(t *testing.T) {
		deps.repository.EXPECT().GetMedicalRecordByID(gomock.Any(), "record-id-1").Return(nil, errors.New("repository error"))

		record, err := service.GetMedicalRecordByID(ctx, "record-id-1")
		require.Error(t, err)
		require.Nil(t, record)
	})
}

func TestService_GetMedicalRecords(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully get all medical records", func(t *testing.T) {
		expected := []domain.MedicalRecord{
			*domain.NewMedicalRecord("account-id-1", "Checkup 1", "First checkup"),
			*domain.NewMedicalRecord("account-id-2", "Checkup 2", "Second checkup"),
		}

		deps.repository.EXPECT().GetMedicalRecords(gomock.Any()).Return(expected, nil)

		records, err := service.GetMedicalRecords(ctx)
		require.NoError(t, err)
		require.Equal(t, expected, records)
	})

	t.Run("should return empty slice when no medical records exist", func(t *testing.T) {
		deps.repository.EXPECT().GetMedicalRecords(gomock.Any()).Return([]domain.MedicalRecord{}, nil)

		records, err := service.GetMedicalRecords(ctx)
		require.NoError(t, err)
		require.Empty(t, records)
	})

	t.Run("should handle error when repository fails", func(t *testing.T) {
		deps.repository.EXPECT().GetMedicalRecords(gomock.Any()).Return(nil, errors.New("repository error"))

		records, err := service.GetMedicalRecords(ctx)
		require.Error(t, err)
		require.Nil(t, records)
	})
}

func TestService_GetMedicalRecordsByAccountID(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully get medical records for account", func(t *testing.T) {
		expected := []domain.MedicalRecord{
			*domain.NewMedicalRecord("account-id-1", "Checkup 1", "First checkup"),
			*domain.NewMedicalRecord("account-id-1", "Checkup 2", "Second checkup"),
		}

		deps.repository.EXPECT().GetMedicalRecordsByAccountID(gomock.Any(), "account-id-1").Return(expected, nil)

		records, err := service.GetMedicalRecordsByAccountID(ctx, "account-id-1")
		require.NoError(t, err)
		require.Equal(t, expected, records)
	})

	t.Run("should return empty slice when no medical records exist for account", func(t *testing.T) {
		deps.repository.EXPECT().GetMedicalRecordsByAccountID(gomock.Any(), "account-id-1").Return([]domain.MedicalRecord{}, nil)

		records, err := service.GetMedicalRecordsByAccountID(ctx, "account-id-1")
		require.NoError(t, err)
		require.Empty(t, records)
	})

	t.Run("should handle error when repository fails", func(t *testing.T) {
		deps.repository.EXPECT().GetMedicalRecordsByAccountID(gomock.Any(), "account-id-1").Return(nil, errors.New("repository error"))

		records, err := service.GetMedicalRecordsByAccountID(ctx, "account-id-1")
		require.Error(t, err)
		require.Nil(t, records)
	})
}

func TestService_UpdateMedicalRecord(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully update medical record", func(t *testing.T) {
		record := domain.NewMedicalRecord("account-id-1", "Checkup", "Annual physical examination")
		record.UpdatedAt = time.Now().UTC()

		deps.repository.EXPECT().UpdateMedicalRecord(gomock.Any(), *record).Return(nil)

		err := service.UpdateMedicalRecord(ctx, *record)
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to update medical record", func(t *testing.T) {
		record := domain.NewMedicalRecord("account-id-1", "Checkup", "Annual physical examination")
		record.UpdatedAt = time.Now().UTC()

		deps.repository.EXPECT().UpdateMedicalRecord(gomock.Any(), *record).Return(errors.New("update error"))

		err := service.UpdateMedicalRecord(ctx, *record)
		require.Error(t, err)
	})
}

func TestService_ArchiveMedicalRecord(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully archive medical record", func(t *testing.T) {
		deps.repository.EXPECT().ArchiveMedicalRecord(gomock.Any(), "record-id-1").Return(nil)

		err := service.ArchiveMedicalRecord(ctx, "record-id-1")
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to archive medical record", func(t *testing.T) {
		deps.repository.EXPECT().ArchiveMedicalRecord(gomock.Any(), "record-id-1").Return(errors.New("archive error"))

		err := service.ArchiveMedicalRecord(ctx, "record-id-1")
		require.Error(t, err)
	})
}
