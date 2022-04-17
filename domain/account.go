package domain

import errs "banking/errors"

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
