package test

import (
	"errors"
	"openhealth/internal/domain"
	"testing"

	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func TestService_CreateHospital(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully create hospital", func(t *testing.T) {
		hospital := domain.NewHospital("St. Test Hospital")

		deps.repository.EXPECT().AddHospital(gomock.Any(), *hospital).Return(nil)

		err := service.CreateHospital(ctx, *hospital)
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to add hospital", func(t *testing.T) {
		hospital := domain.NewHospital("St. Test Hospital")

		deps.repository.EXPECT().AddHospital(gomock.Any(), *hospital).Return(errors.New("add error"))

		err := service.CreateHospital(ctx, *hospital)
		require.Error(t, err)
	})
}

func TestService_GetHospitalByID(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully get hospital", func(t *testing.T) {
		hospital := domain.NewHospital("St. Test Hospital")

		deps.repository.EXPECT().GetHospitalByID(gomock.Any(), "hospital-id-1").Return(hospital, nil)

		result, err := service.GetHospitalByID(ctx, "hospital-id-1")
		require.NoError(t, err)
		require.Equal(t, hospital, result)
	})

	t.Run("should handle error when hospital is not found", func(t *testing.T) {
		deps.repository.EXPECT().GetHospitalByID(gomock.Any(), "hospital-id-1").Return(nil, nil)

		hospital, err := service.GetHospitalByID(ctx, "hospital-id-1")

		require.ErrorIs(t, err, domain.HospitalNotFoundError)
		require.Nil(t, hospital)
	})

	t.Run("should handle error when repository fails to get hospital", func(t *testing.T) {
		deps.repository.EXPECT().GetHospitalByID(gomock.Any(), "hospital-id-1").Return(nil, errors.New("repository error"))

		hospital, err := service.GetHospitalByID(ctx, "hospital-id-1")

		require.Error(t, err)
		require.Nil(t, hospital)
	})
}

func TestService_GetAllHospitals(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully return hospitals up to the limit", func(t *testing.T) {
		expected := []domain.Hospital{
			*domain.NewHospital("General Hospital"),
			*domain.NewHospital("Regional Hospital"),
		}

		deps.repository.EXPECT().GetHospitals(gomock.Any(), 2).Return(expected, nil)

		hospitals, err := service.GetHospitals(ctx, 2)
		require.NoError(t, err)
		require.Len(t, hospitals, 2)
		require.Equal(t, expected, hospitals)
	})

	t.Run("should return empty slice when no hospitals exist", func(t *testing.T) {
		deps.repository.EXPECT().GetHospitals(gomock.Any(), 5).Return([]domain.Hospital{}, nil)

		hospitals, err := service.GetHospitals(ctx, 5)
		require.NoError(t, err)
		require.Empty(t, hospitals)
	})

	t.Run("should handle error when repository fails", func(t *testing.T) {
		deps.repository.EXPECT().GetHospitals(gomock.Any(), 10).Return(nil, errors.New("repository error"))

		hospitals, err := service.GetHospitals(ctx, 10)
		require.Error(t, err)
		require.Nil(t, hospitals)
	})
}

func TestService_UpdateHospital(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully update hospital", func(t *testing.T) {
		hospital := domain.NewHospital("St. Test Hospital")
		hospital.Name = "Updated Hospital Name"

		deps.repository.EXPECT().UpdateHospital(gomock.Any(), *hospital).Return(nil)

		err := service.UpdateHospital(ctx, *hospital)
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to update hospital", func(t *testing.T) {
		hospital := domain.NewHospital("St. Test Hospital")

		deps.repository.EXPECT().UpdateHospital(gomock.Any(), *hospital).Return(errors.New("update error"))

		err := service.UpdateHospital(ctx, *hospital)
		require.Error(t, err)
	})
}

func TestService_ArchiveHospital(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully archive hospital", func(t *testing.T) {
		deps.repository.EXPECT().ArchiveHospital(gomock.Any(), "hospital-id-1").Return(nil)

		err := service.ArchiveHospital(ctx, "hospital-id-1")
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to archive hospital", func(t *testing.T) {
		deps.repository.EXPECT().ArchiveHospital(gomock.Any(), "hospital-id-1").Return(errors.New("archive error"))

		err := service.ArchiveHospital(ctx, "hospital-id-1")
		require.Error(t, err)
	})
}

func TestService_AddHospitalToAccountAssignment(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully add hospital assignment", func(t *testing.T) {
		deps.repository.EXPECT().AddHospitalAssignment(gomock.Any(), gomock.Any()).Return(nil)

		err := service.AddHospitalToAccountAssignment(ctx, "account-id-1", "hospital-id-1")
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to add hospital assignment", func(t *testing.T) {
		deps.repository.EXPECT().AddHospitalAssignment(gomock.Any(), gomock.Any()).Return(errors.New("assign error"))

		err := service.AddHospitalToAccountAssignment(ctx, "account-id-1", "hospital-id-1")
		require.Error(t, err)
	})
}

func TestService_RemoveHospitalToAccountAssignment(t *testing.T) {
	ctx := t.Context()

	service, deps := newMockService(t)

	t.Run("should successfully remove hospital to account assignment", func(t *testing.T) {
		deps.repository.EXPECT().RemoveHospitalToAccountAssignment(gomock.Any(), "assignment-id-1").Return(nil)

		err := service.RemoveHospitalToAccountAssignment(ctx, "assignment-id-1")
		require.NoError(t, err)
	})

	t.Run("should handle error when repository fails to remove hospital to account assignment", func(t *testing.T) {
		deps.repository.EXPECT().RemoveHospitalToAccountAssignment(gomock.Any(), "assignment-id-1").Return(errors.New("remove error"))

		err := service.RemoveHospitalToAccountAssignment(ctx, "assignment-id-1")
		require.Error(t, err)
	})
}
