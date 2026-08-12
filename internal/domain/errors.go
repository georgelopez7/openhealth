package domain

import "errors"

var (
	AccountNotFoundError = errors.New("account not found")
	DoctorNotFoundError  = errors.New("doctor not found")
)
