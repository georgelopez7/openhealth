package repository

import (
	"context"
	"database/sql"

	"openhealth/internal/domain"
	"openhealth/internal/pkg/postgres"
)

// AddAccount - adds a new account to the database
func (r *Repository) AddAccount(ctx context.Context, account domain.Account) error {
	dx := postgres.GetTxOrDB(ctx, r.db)

	_, err := dx.ExecContext(ctx, `
		INSERT INTO accounts (id, first_name, last_name, age, email, avatar, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, account.ID, account.FirstName, account.LastName, account.Age, account.Email, account.Avatar, account.Status, account.CreatedAt, account.UpdatedAt)

	return err
}

// GetAccountByID - get an account by its ID
func (r *Repository) GetAccountByID(ctx context.Context, id string) (*domain.Account, error) {
	var account domain.Account

	err := r.db.GetContext(ctx, &account, `
		SELECT id, first_name, last_name, age, email, avatar, status, created_at, updated_at
		FROM accounts
		WHERE id = $1
	`, id)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &account, nil
	default:
		return nil, err
	}
}

// GetAccounts - get accounts up to the provided limit
func (r *Repository) GetAccounts(ctx context.Context, limit int) ([]domain.Account, error) {
	var accounts []domain.Account

	err := r.db.SelectContext(ctx, &accounts, `
		SELECT id, first_name, last_name, age, email, avatar, status, created_at, updated_at
		FROM accounts
		LIMIT $1
	`, limit)

	return accounts, err
}

// UpdateAccount - updates an existing account
func (r *Repository) UpdateAccount(ctx context.Context, account domain.Account) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE accounts
		SET first_name = $1, last_name = $2, age = $3, email = $4, avatar = $5, status = $6, updated_at = $7
		WHERE id = $8
	`, account.FirstName, account.LastName, account.Age, account.Email, account.Avatar, account.Status, account.UpdatedAt, account.ID)

	return err
}

// ArchiveAccount - archives an existing account
func (r *Repository) ArchiveAccount(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE accounts
		SET status = $1, updated_at = NOW()
		WHERE id = $2
	`, domain.AccountStatusArchived, id)

	return err
}

// GetDoctorIDByAccountID - get the currently assigned doctor ID for an account
func (r *Repository) GetDoctorIDByAccountID(ctx context.Context, accountID string) (string, error) {
	var doctorID string

	err := r.db.GetContext(ctx, &doctorID, `
		SELECT doctor_id
		FROM doctor_to_account_assignments
		WHERE account_id = $1
		AND valid_to IS NULL
		ORDER BY valid_from DESC
		LIMIT 1
	`, accountID)

	switch err {
	case sql.ErrNoRows:
		return "", nil
	case nil:
		return doctorID, nil
	default:
		return "", err
	}
}

// ResetAccounts - resets all accounts
func (r *Repository) ResetAccounts(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
		TRUNCATE TABLE accounts CASCADE
	`)

	return err
}
