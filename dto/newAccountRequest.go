package dto

import (
	"strings"

	"github.com/saskaradit/finance-app/errs"
)

type NewAccountRequest struct {
	CustomerId string  `json:"customer_id"`
	Type       string  `json:"type"`
	Balance    float64 `json:"balance"`
}

func (n NewAccountRequest) Validate() *errs.AppError {
	if n.Balance < 500 {
		return errs.NewValidationError("To open you need to deposit at least 500")
	}
	if strings.ToLower(n.Type) != "saving" && strings.ToLower(n.Type) != "checking" {
		return errs.NewValidationError("Account should be type savings or checkings")
	}
	return nil
}
