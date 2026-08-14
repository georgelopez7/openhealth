package relay

import (
	"context"

	"openhealth/internal/event"
)

func (r *Relay) handleDoctorCreated(ctx context.Context, evt event.Event) error {
	var p event.DoctorCreatedEventPayload
	if err := processEventPayload(evt.Payload, &p); err != nil {
		return err
	}

	return r.authorizationService.AddDoctorAccount(ctx, p.Doctor.ID, p.Doctor.AccountID)
}

func (r *Relay) handleDoctorDeleted(ctx context.Context, evt event.Event) error {
	var p event.DoctorDeletedEventPayload
	if err := processEventPayload(evt.Payload, &p); err != nil {
		return err
	}

	return r.authorizationService.RemoveDoctorAccount(ctx, p.Doctor.ID, p.Doctor.AccountID)
}

func (r *Relay) handleNurseCreated(ctx context.Context, evt event.Event) error {
	var p event.NurseCreatedEventPayload
	if err := processEventPayload(evt.Payload, &p); err != nil {
		return err
	}

	return r.authorizationService.AddNurseAccount(ctx, p.Nurse.ID, p.Nurse.AccountID)
}

func (r *Relay) handleNurseDeleted(ctx context.Context, evt event.Event) error {
	var p event.NurseDeletedEventPayload
	if err := processEventPayload(evt.Payload, &p); err != nil {
		return err
	}

	return r.authorizationService.RemoveNurseAccount(ctx, p.Nurse.ID, p.Nurse.AccountID)
}

func (r *Relay) handleDoctorToAccountAssignmentCreated(ctx context.Context, evt event.Event) error {
	var p event.DoctorToAccountAssignmentCreatedEventPayload
	if err := processEventPayload(evt.Payload, &p); err != nil {
		return err
	}

	return r.authorizationService.AddDoctorToAccountAssignment(ctx, p.Assignment.AccountID, p.Assignment.DoctorID)
}

func (r *Relay) handleDoctorToAccountAssignmentRemoved(ctx context.Context, evt event.Event) error {
	var p event.DoctorToAccountAssignmentRemovedEventPayload
	if err := processEventPayload(evt.Payload, &p); err != nil {
		return err
	}

	return r.authorizationService.RemoveDoctorToAccountAssignment(ctx, p.Assignment.AccountID, p.Assignment.DoctorID)
}

func (r *Relay) handleHospitalToAccountAssignmentCreated(ctx context.Context, evt event.Event) error {
	var p event.HospitalToAccountAssignmentCreatedEventPayload
	if err := processEventPayload(evt.Payload, &p); err != nil {
		return err
	}

	return r.authorizationService.AddHospitalToAccountAssignment(ctx, p.Assignment.AccountID, p.Assignment.HospitalID)
}

func (r *Relay) handleHospitalToAccountAssignmentRemoved(ctx context.Context, evt event.Event) error {
	var p event.HospitalToAccountAssignmentRemovedEventPayload
	if err := processEventPayload(evt.Payload, &p); err != nil {
		return err
	}

	return r.authorizationService.RemoveHospitalToAccountAssignment(ctx, p.Assignment.AccountID, p.Assignment.HospitalID)
}

func (r *Relay) handleNurseToHospitalAssignmentCreated(ctx context.Context, evt event.Event) error {
	var p event.NurseToHospitalAssignmentCreatedEventPayload
	if err := processEventPayload(evt.Payload, &p); err != nil {
		return err
	}

	return r.authorizationService.AddNurseToHospitalAssignment(ctx, p.Assignment.NurseID, p.Assignment.HospitalID)
}

func (r *Relay) handleNurseToHospitalAssignmentRemoved(ctx context.Context, evt event.Event) error {
	var p event.NurseToHospitalAssignmentRemovedEventPayload
	if err := processEventPayload(evt.Payload, &p); err != nil {
		return err
	}

	return r.authorizationService.RemoveNurseToHospitalAssignment(ctx, p.Assignment.NurseID, p.Assignment.HospitalID)
}

func (r *Relay) handleMedicalRecordCreated(ctx context.Context, evt event.Event) error {
	var p event.MedicalRecordCreatedEventPayload
	if err := processEventPayload(evt.Payload, &p); err != nil {
		return err
	}

	return r.authorizationService.AddMedicalRecordOwner(ctx, p.Record.ID, p.Record.AccountID)
}

func (r *Relay) handleMedicalRecordDeleted(ctx context.Context, evt event.Event) error {
	var p event.MedicalRecordDeletedEventPayload
	if err := processEventPayload(evt.Payload, &p); err != nil {
		return err
	}

	return r.authorizationService.RemoveMedicalRecordOwner(ctx, p.RecordID, p.AccountID)
}
