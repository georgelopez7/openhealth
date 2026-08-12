package test

import (
	"context"
	"errors"
	"openhealth/internal/domain"
	"testing"

	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func TestService_CreateNurse(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully create nurse", func(t *testing.T) {
		nurse := domain.NewNurse("account-id-1")

		deps.repository.EXPECT().AddNurse(gomock.Any(), *nurse).Return(nil)

		err := service.CreateNurse(ctx, *nurse)
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to add nurse", func(t *testing.T) {
		nurse := domain.NewNurse("account-id-1")

		deps.repository.EXPECT().AddNurse(gomock.Any(), *nurse).Return(errors.New("add error"))

		err := service.CreateNurse(ctx, *nurse)
		require.Error(t, err)
	})
}

func TestService_GetNurseByID(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully get nurse", func(t *testing.T) {
		nurse := domain.NewNurse("account-id-1")

		deps.repository.EXPECT().GetNurseByID(gomock.Any(), "nurse-id-1").Return(nurse, nil)

		result, err := service.GetNurseByID(ctx, "nurse-id-1")
		require.NoError(t, err)
		require.Equal(t, nurse, result)
	})

	t.Run("should handle error when nurse is not found", func(t *testing.T) {
		deps.repository.EXPECT().GetNurseByID(gomock.Any(), "nurse-id-1").Return(nil, nil)

		nurse, err := service.GetNurseByID(ctx, "nurse-id-1")

		require.ErrorIs(t, err, domain.NurseNotFoundError)
		require.Nil(t, nurse)
	})

	t.Run("should handle error when repository fails to get nurse", func(t *testing.T) {
		deps.repository.EXPECT().GetNurseByID(gomock.Any(), "nurse-id-1").Return(nil, errors.New("repository error"))

		nurse, err := service.GetNurseByID(ctx, "nurse-id-1")

		require.Error(t, err)
		require.Nil(t, nurse)
	})
}

func TestService_GetAllNurses(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully return nurses up to the limit", func(t *testing.T) {
		expected := []domain.Nurse{
			*domain.NewNurse("account-id-1"),
			*domain.NewNurse("account-id-2"),
		}

		deps.repository.EXPECT().GetNurses(gomock.Any(), 2).Return(expected, nil)

		nurses, err := service.GetNurses(ctx, 2)
		require.NoError(t, err)
		require.Len(t, nurses, 2)
		require.Equal(t, expected, nurses)
	})

	t.Run("should return empty slice when no nurses exist", func(t *testing.T) {
		deps.repository.EXPECT().GetNurses(gomock.Any(), 5).Return([]domain.Nurse{}, nil)

		nurses, err := service.GetNurses(ctx, 5)
		require.NoError(t, err)
		require.Empty(t, nurses)
	})

	t.Run("should handle error when repository fails", func(t *testing.T) {
		deps.repository.EXPECT().GetNurses(gomock.Any(), 10).Return(nil, errors.New("repository error"))

		nurses, err := service.GetNurses(ctx, 10)
		require.Error(t, err)
		require.Nil(t, nurses)
	})
}

func TestService_UpdateNurseStatusByID(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully archive nurse and expire assignments", func(t *testing.T) {
		deps.tx.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(context.Context) error) error {
				return fn(ctx)
			},
		)

		deps.repository.EXPECT().UpdateNurseStatusByID(gomock.Any(), "nurse-id-1", domain.NurseStatusArchived).Return(nil)
		deps.repository.EXPECT().RemoveAllNurseToHospitalAssignments(gomock.Any(), "nurse-id-1").Return(nil)

		err := service.UpdateNurseStatusByID(ctx, "nurse-id-1", domain.NurseStatusArchived)
		require.NoError(t, err)
	})

	t.Run("should not expire assignments when status is not archived", func(t *testing.T) {
		deps.tx.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(context.Context) error) error {
				return fn(ctx)
			},
		)

		deps.repository.EXPECT().UpdateNurseStatusByID(gomock.Any(), "nurse-id-1", domain.NurseStatusActive).Return(nil)

		err := service.UpdateNurseStatusByID(ctx, "nurse-id-1", domain.NurseStatusActive)
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to update nurse status", func(t *testing.T) {
		deps.tx.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(context.Context) error) error {
				return fn(ctx)
			},
		)

		deps.repository.EXPECT().UpdateNurseStatusByID(gomock.Any(), "nurse-id-1", domain.NurseStatusArchived).Return(errors.New("update error"))

		err := service.UpdateNurseStatusByID(ctx, "nurse-id-1", domain.NurseStatusArchived)
		require.Error(t, err)
	})

	t.Run("should handle error when repository fails to expire assignments", func(t *testing.T) {
		deps.tx.EXPECT().WithTransaction(gomock.Any(), gomock.Any()).DoAndReturn(
			func(ctx context.Context, fn func(context.Context) error) error {
				return fn(ctx)
			},
		)

		deps.repository.EXPECT().UpdateNurseStatusByID(gomock.Any(), "nurse-id-1", domain.NurseStatusArchived).Return(nil)
		deps.repository.EXPECT().RemoveAllNurseToHospitalAssignments(gomock.Any(), "nurse-id-1").Return(errors.New("remove error"))

		err := service.UpdateNurseStatusByID(ctx, "nurse-id-1", domain.NurseStatusArchived)
		require.Error(t, err)
	})
}

func TestService_AddNurseToHospitalAssignment(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully add nurse to hospital assignment", func(t *testing.T) {
		deps.repository.EXPECT().AddNurseToHospitalAssignment(gomock.Any(), gomock.Any()).Return(nil)

		err := service.AddNurseToHospitalAssignment(ctx, "nurse-id-1", "hospital-id-1")
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to add nurse to hospital assignment", func(t *testing.T) {
		deps.repository.EXPECT().AddNurseToHospitalAssignment(gomock.Any(), gomock.Any()).Return(errors.New("assign error"))

		err := service.AddNurseToHospitalAssignment(ctx, "nurse-id-1", "hospital-id-1")
		require.Error(t, err)
	})
}
