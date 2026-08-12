package http

import "openhealth/internal/domain"

type CreateAccountInput struct {
	Body struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Age       int    `json:"age"`
		Email     string `json:"email"`
	}
}

type CreateAccountResponse struct {
	Body struct {
		Account domain.Account `json:"account"`
	}
}

type UpdateAccountInput struct {
	AccountID string `path:"accountID"`
	Body struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Age       int    `json:"age"`
		Email     string `json:"email"`
	}
}

type UpdateAccountResponse struct{}

type GetAccountByIDInput struct {
	AccountID string `path:"accountID"`
}

type GetAccountByIDResponse struct {
	Body struct {
		Account domain.Account `json:"account"`
	}
}

type ListAccountsInput struct {
	Limit int `query:"limit" default:"10"`
}

type ListAccountsResponse struct {
	Body struct {
		Accounts []domain.Account `json:"accounts"`
	}
}

type ArchiveAccountInput struct {
	AccountID string `path:"accountID"`
}

type ArchiveAccountResponse struct{}
