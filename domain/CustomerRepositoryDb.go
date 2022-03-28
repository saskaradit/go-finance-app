package domain

import (
	"database/sql"
	"time"

	"github.com/saskaradit/finance-app/errs"
	"github.com/saskaradit/finance-app/logger"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var findAllSql string
	var rows *sql.Rows
	var err error
	if status == "" {
		findAllSql = "select customer_id, name, city,zipcode,date_of_birth, status from customers"
		rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql = "select customer_id, name, city,zipcode,date_of_birth, status from customers where status = ?"
		rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		errs.NewUnexpectedError("unexpected database error")
		logger.Error("Error while getting customers" + err.Error())
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.DateOfBirth, &c.Status)
		if err != nil {
			logger.Error("Error while getting customers" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city,zipcode,date_of_birth, status from customers where customer_id = ?"
	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while getting customers" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
