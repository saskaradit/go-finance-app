package service

import (
	"time"

	"github.com/saskaradit/finance-app/domain"
	"github.com/saskaradit/finance-app/dto"
	"github.com/saskaradit/finance-app/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	account := domain.Account{
		AccountId:  "",
		CustomerId: req.CustomerId,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		Type:       req.Type,
		Balance:    req.Balance,
		Status:     "1",
	}
	newAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}
	res := newAccount.ToNewAccountResponseDto()
	return &res, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
