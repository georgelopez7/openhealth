package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	AccountStatusActive   = "active"
	AccountStatusArchived = "archived"
)

type Account struct {
	ID        string    `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Age       int       `json:"age" db:"age"`
	Email     string    `json:"email" db:"email"`
	Avatar     string    `json:"avatar" db:"avatar"`
	Status     string    `json:"status" db:"status"`
	DoctorID   string    `json:"doctor_id" db:"-"`
	HospitalID string    `json:"hospital_id" db:"-"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewAccount(firstName string, lastName string, age int, email string, avatar string) *Account {
	now := time.Now().UTC()
	return &Account{
		ID:        uuid.NewString(),
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Email:     email,
		Avatar:    avatar,
		Status:    AccountStatusActive,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
