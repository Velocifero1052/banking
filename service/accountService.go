package service

import (
	"banking/domain"
	"banking/dto"
	errs "banking/errors"
	"fmt"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repository domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	acc := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	fmt.Println(acc)
	newAccount, err := s.repository.Save(acc)

	if err != nil {
		return nil, err
	}

	resp := newAccount.ToNewAccountResponseDto()

	return &resp, nil
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repository: repository}
}
