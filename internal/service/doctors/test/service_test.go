package test

import (
	"context"
	"errors"
	"openhealth/internal/domain"
	"testing"

	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func TestService_CreateDoctor(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully create doctor", func(t *testing.T) {
		doctor := domain.NewDoctor("account-id-1")

		deps.repository.EXPECT().AddDoctor(gomock.Any(), *doctor).Return(nil)

		err := service.CreateDoctor(ctx, *doctor)
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to add doctor", func(t *testing.T) {
		doctor := domain.NewDoctor("account-id-1")

		deps.repository.EXPECT().AddDoctor(gomock.Any(), *doctor).Return(errors.New("add error"))

		err := service.CreateDoctor(ctx, *doctor)
		require.Error(t, err)
	})
}

func TestService_GetDoctorByID(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully get doctor", func(t *testing.T) {
		doctor := domain.NewDoctor("account-id-1")

		deps.repository.EXPECT().GetDoctorByID(gomock.Any(), "doctor-id-1").Return(doctor, nil)

		result, err := service.GetDoctorByID(ctx, "doctor-id-1")
		require.NoError(t, err)
		require.Equal(t, doctor, result)
	})

	t.Run("should handle error when doctor is not found", func(t *testing.T) {
		deps.repository.EXPECT().GetDoctorByID(gomock.Any(), "doctor-id-1").Return(nil, nil)

		doctor, err := service.GetDoctorByID(ctx, "doctor-id-1")

		require.ErrorIs(t, err, domain.DoctorNotFoundError)
		require.Nil(t, doctor)
	})

	t.Run("should handle error when repository fails to get doctor", func(t *testing.T) {
		deps.repository.EXPECT().GetDoctorByID(gomock.Any(), "doctor-id-1").Return(nil, errors.New("repository error"))

		doctor, err := service.GetDoctorByID(ctx, "doctor-id-1")

		require.Error(t, err)
		require.Nil(t, doctor)
	})
}

func TestService_GetAllDoctors(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully return doctors up to the limit", func(t *testing.T) {
		expected := []domain.Doctor{
			*domain.NewDoctor("account-id-1"),
			*domain.NewDoctor("account-id-2"),
		}

		deps.repository.EXPECT().GetDoctors(gomock.Any(), 2).Return(expected, nil)

		doctors, err := service.GetDoctors(ctx, 2)
		require.NoError(t, err)
		require.Len(t, doctors, 2)
		require.Equal(t, expected, doctors)
	})

	t.Run("should return empty slice when no doctors exist", func(t *testing.T) {
		deps.repository.EXPECT().GetDoctors(gomock.Any(), 5).Return([]domain.Doctor{}, nil)

		doctors, err := service.GetDoctors(ctx, 5)
		require.NoError(t, err)
		require.Empty(t, doctors)
	})

	t.Run("should handle error when repository fails", func(t *testing.T) {
		deps.repository.EXPECT().GetDoctors(gomock.Any(), 10).Return(nil, errors.New("repository error"))

		doctors, err := service.GetDoctors(ctx, 10)
		require.Error(t, err)
		require.Nil(t, doctors)
	})
}

func TestService_UpdateDoctorStatusByID(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully archive doctor and expire assignments", func(t *testing.T) {
		deps.tx.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(context.Context) error) error {
				return fn(ctx)
			},
		)

		deps.repository.EXPECT().UpdateDoctorStatusByID(gomock.Any(), "doctor-id-1", domain.DoctorStatusArchived).Return(nil)
		deps.repository.EXPECT().RemoveAllDoctorToAccountAssignments(gomock.Any(), "doctor-id-1").Return(nil)

		err := service.UpdateDoctorStatusByID(ctx, "doctor-id-1", domain.DoctorStatusArchived)
		require.NoError(t, err)
	})

	t.Run("should not expire assignments when status is not archived", func(t *testing.T) {
		deps.tx.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(context.Context) error) error {
				return fn(ctx)
			},
		)

		deps.repository.EXPECT().UpdateDoctorStatusByID(gomock.Any(), "doctor-id-1", domain.DoctorStatusActive).Return(nil)

		err := service.UpdateDoctorStatusByID(ctx, "doctor-id-1", domain.DoctorStatusActive)
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to update doctor status", func(t *testing.T) {
		deps.tx.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(context.Context) error) error {
				return fn(ctx)
			},
		)

		deps.repository.EXPECT().UpdateDoctorStatusByID(gomock.Any(), "doctor-id-1", domain.DoctorStatusArchived).Return(errors.New("update error"))

		err := service.UpdateDoctorStatusByID(ctx, "doctor-id-1", domain.DoctorStatusArchived)
		require.Error(t, err)
	})

	t.Run("should handle error when repository fails to expire assignments", func(t *testing.T) {
		deps.tx.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(context.Context) error) error {
				return fn(ctx)
			},
		)

		deps.repository.EXPECT().UpdateDoctorStatusByID(gomock.Any(), "doctor-id-1", domain.DoctorStatusArchived).Return(nil)
		deps.repository.EXPECT().RemoveAllDoctorToAccountAssignments(gomock.Any(), "doctor-id-1").Return(errors.New("remove error"))

		err := service.UpdateDoctorStatusByID(ctx, "doctor-id-1", domain.DoctorStatusArchived)
		require.Error(t, err)
	})
}

func TestService_AddDoctorToAccountAssignment(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully add doctor assignment", func(t *testing.T) {
		deps.repository.EXPECT().AddDoctorToAccountAssignment(gomock.Any(), gomock.Any()).Return(nil)

		assignment, err := service.AddDoctorToAccountAssignment(ctx, "account-id-1", "doctor-id-1")
		require.NoError(t, err)
		require.Equal(t, "account-id-1", assignment.AccountID)
		require.Equal(t, "doctor-id-1", assignment.DoctorID)
		require.NotEmpty(t, assignment.ID)
	})

	t.Run("should handle error when repository fails to add doctor assignment", func(t *testing.T) {
		deps.repository.EXPECT().AddDoctorToAccountAssignment(gomock.Any(), gomock.Any()).Return(errors.New("assign error"))

		assignment, err := service.AddDoctorToAccountAssignment(ctx, "account-id-1", "doctor-id-1")
		require.Error(t, err)
		require.Nil(t, assignment)
	})
}
