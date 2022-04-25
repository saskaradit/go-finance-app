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

const dbTSLayout = "2006-01-02 15:04:05"

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountId: a.AccountId}
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/ashishjuyal/banking/domain AccountRepository
type AccountRepository interface {
	Save(account Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Balance >= amount
}

func NewAccount(customerId, accountType string, amount float64) Account {
	return Account{
		CustomerId: customerId,
		CreatedAt:  dbTSLayout,
		Type:       accountType,
		Balance:    amount,
		Status:     "1",
	}
}
