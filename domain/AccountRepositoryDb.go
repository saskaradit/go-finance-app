package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/saskaradit/finance-app/errs"
	"github.com/saskaradit/finance-app/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT into accounts (customer_id, createdAt, account_type, balance, status) VALUES (?,?,?,?,?)"
	res, err := d.client.Exec(sqlInsert, a.CustomerId, a.CreatedAt, a.Type, a.Balance, a.Status)
	if err != nil {
		logger.Error("Error while creating a new account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last inserted ID" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
