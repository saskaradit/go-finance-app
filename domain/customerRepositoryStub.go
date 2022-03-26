package domain

import "github.com/saskaradit/finance-app/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Rad", "JKT", "14045", "20-05-1996", "Active"},
		{"2", "Hehe", "JKT", "14045", "20-05-1995", "Active"},
		{"101", "Jake", "JKT", "14045", "20-05-1994", "Active"},
	}

	return CustomerRepositoryStub{customers}
}
