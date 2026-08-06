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
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewAccount(firstName string, lastName string, age int, email string) *Account {
	now := time.Now().UTC()
	return &Account{
		ID:        uuid.NewString(),
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Email:     email,
		Status:    AccountStatusActive,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
