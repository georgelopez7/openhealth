package domain

import "errors"

var (
	AccountNotFoundError  = errors.New("account not found")
	DoctorNotFoundError   = errors.New("doctor not found")
	NurseNotFoundError    = errors.New("nurse not found")
	HospitalNotFoundError = errors.New("hospital not found")
)
