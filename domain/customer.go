package domain

import (
	"github.com/saskaradit/finance-app/dto"
	"github.com/saskaradit/finance-app/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	//  status == 1, status == 0, status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

func (c Customer) StatusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      statusAsText,
	}
}
