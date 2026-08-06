package http

import (
	"context"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"openhealth/internal/domain"
)

func (s *Server) CreateAccountHandler(ctx context.Context, input *CreateAccountInput) (*CreateAccountResponse, error) {
	account := domain.NewAccount(
		input.Body.FirstName,
		input.Body.LastName,
		input.Body.Age,
		input.Body.Email,
	)

	err := s.AccountSVC.CreateAccount(ctx, *account)
	if err != nil {
		return nil, err
	}

	resp := &CreateAccountResponse{}
	resp.Body.Account = *account

	return resp, nil
}

func (s *Server) UpdateAccountHandler(ctx context.Context, input *UpdateAccountInput) (*UpdateAccountResponse, error) {
	account := domain.Account{
		ID:        input.ID,
		FirstName: input.Body.FirstName,
		LastName:  input.Body.LastName,
		Age:       input.Body.Age,
		Email:     input.Body.Email,
		UpdatedAt: time.Now().UTC(),
	}

	err := s.AccountSVC.UpdateAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	return &UpdateAccountResponse{}, nil
}

func (s *Server) GetAccountByIDHandler(ctx context.Context, input *GetAccountByIDInput) (*GetAccountByIDResponse, error) {
	account, err := s.AccountSVC.GetAccountByID(ctx, input.ID)
	switch err {
	case nil:
		resp := &GetAccountByIDResponse{}
		resp.Body.Account = *account
		return resp, nil
	case domain.AccountNotFoundError:
		return nil, huma.Error404NotFound("account not found")
	default:
		return nil, err
	}
}

func (s *Server) ListAccountsHandler(ctx context.Context, input *ListAccountsInput) (*ListAccountsResponse, error) {
	accounts, err := s.AccountSVC.GetAccounts(ctx, input.Limit)
	if err != nil {
		return nil, err
	}

	resp := &ListAccountsResponse{}
	resp.Body.Accounts = accounts

	return resp, nil
}
