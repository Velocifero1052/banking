package domain

import (
	errs "banking/errors"
	"banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) save(account Account) (*Account, *errs.AppError) {
	sqlInsert := "insert into accounts (customer_id, opening_date, account_type, amount, status) values ($1, $2, $3, $4, $5)"
	result, err := d.client.Exec(sqlInsert, account.CustomerId, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logger.Error("error while creating new account: " + err.Error())
		return nil, errs.NewInternalServerError("error while creating new account")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting last inserted id: " + err.Error())
		return nil, errs.NewInternalServerError("error while getting last inserted id")
	}

	account.AccountId = strconv.FormatInt(id, 10)
	return &account, nil
}

func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client}
}
