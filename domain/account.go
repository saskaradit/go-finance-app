package domain

import (
	"github.com/saskaradit/finance-app/dto"
	"github.com/saskaradit/finance-app/errs"
)

type Account struct {
	AccountId  string
	CustomerId string
	CreatedAt  string
	Type       string
	Balance    float64
	Status     string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}
