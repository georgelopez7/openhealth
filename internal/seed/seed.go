package seed

import (
	"context"
	"fmt"
	"log/slog"

	"openhealth/internal/domain"
)

type Services struct {
	accountSVC       AccountService
	doctorSVC        DoctorService
	nurseSVC         NurseService
	hospitalSVC      HospitalService
	medrecordSVC     MedicalRecordService
}

func NewServices(accountSVC AccountService, doctorSVC DoctorService, nurseSVC NurseService, hospitalSVC HospitalService, medrecordSVC MedicalRecordService) Services {
	return Services{
		accountSVC:   accountSVC,
		doctorSVC:    doctorSVC,
		nurseSVC:     nurseSVC,
		hospitalSVC:  hospitalSVC,
		medrecordSVC: medrecordSVC,
	}
}

// Seed initializes the demo dataset only when the database has no accounts.
func Seed(ctx context.Context, services Services) error {
	existingAccounts, err := services.accountSVC.GetAccounts(ctx, 1)
	if err != nil {
		return fmt.Errorf("failed to check for existing accounts: %w", err)
	}

	if len(existingAccounts) > 0 {
		slog.Info("skipping database seed - accounts already exist")
		return nil
	}

	ash := domain.NewAccount("Ash", "Ketchum", 10, "ash.ketchum@pokemon.com", "ash-ketchum")
	misty := domain.NewAccount("Misty", "Waterflower", 12, "misty.waterflower@pokemon.com", "misty-waterflower")
	oak := domain.NewAccount("Professor", "Oak", 65, "professor.oak@pokemon.com", "professor-oak")
	joy := domain.NewAccount("Nurse", "Joy", 24, "nurse.joy@pokemon.com", "nurse-joy")
	brock := domain.NewAccount("Brock", "Harrison", 15, "brock.harrison@pokemon.com", "brock-harrison")

	accounts := []*domain.Account{ash, misty, oak, joy, brock}
	for _, account := range accounts {
		if err := services.accountSVC.CreateAccount(ctx, *account); err != nil {
			return fmt.Errorf("failed to create account %s %s: %w", account.FirstName, account.LastName, err)
		}
	}

	doctor := domain.NewDoctor(oak.ID)
	if err := services.doctorSVC.CreateDoctor(ctx, *doctor); err != nil {
		return fmt.Errorf("failed to create doctor: %w", err)
	}

	if _, err := services.doctorSVC.AddDoctorToAccountAssignment(ctx, ash.ID, doctor.ID); err != nil {
		return fmt.Errorf("failed to assign doctor to Ash: %w", err)
	}

	if _, err := services.doctorSVC.AddDoctorToAccountAssignment(ctx, brock.ID, doctor.ID); err != nil {
		return fmt.Errorf("failed to assign doctor to Brock: %w", err)
	}

	nurse := domain.NewNurse(joy.ID)
	if err := services.nurseSVC.CreateNurse(ctx, *nurse); err != nil {
		return fmt.Errorf("failed to create nurse: %w", err)
	}

	hospital := domain.NewHospital("Pokemon Center")
	if err := services.hospitalSVC.CreateHospital(ctx, *hospital); err != nil {
		return fmt.Errorf("failed to create hospital: %w", err)
	}

	if _, err := services.nurseSVC.AddNurseToHospitalAssignment(ctx, nurse.ID, hospital.ID); err != nil {
		return fmt.Errorf("failed to assign Nurse Joy to Pokemon Center: %w", err)
	}

	if _, err := services.hospitalSVC.AddHospitalToAccountAssignment(ctx, ash.ID, hospital.ID); err != nil {
		return fmt.Errorf("failed to assign Pokemon Center to Ash: %w", err)
	}

	records := []*domain.MedicalRecord{
		domain.NewMedicalRecord(ash.ID, "Pikachu Thunder Shock Incident", "Treatment for Ash Ketchum after Pikachu's Thunder Shock triggered his nosebleed allergy."),
		domain.NewMedicalRecord(misty.ID, "Tentacool Beach Sting", "Treatment for Misty Waterflower after being stung during the Tentacool invasion of Porta Vista."),
		domain.NewMedicalRecord(misty.ID, "Misty Psyduck Headache Incident", "Consultation for Misty Waterflower after enduring Psyduck's constant headaches."),
		domain.NewMedicalRecord(brock.ID, "Onix Training Injury", "Treatment for Brock Harrison after being crushed while training with Onix in the Pewter City Gym."),
		domain.NewMedicalRecord(brock.ID, "Vulpix Flame Mishap", "Treatment for minor burns Brock Harrison sustained while feeding Vulpix."),
	}

	for _, record := range records {
		if err := services.medrecordSVC.CreateMedicalRecord(ctx, *record); err != nil {
			return fmt.Errorf("failed to create medical record %q: %w", record.Title, err)
		}
	}

	slog.Info("database demo seed completed")

	return nil
}
