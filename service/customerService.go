package service

import (
	"banking/domain"
	"banking/dto"
	errs "banking/errors"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repository.FindAll(status)

	if err != nil {
		return nil, err
	}

	resp := make([]dto.CustomerResponse, 0)

	for i := 0; i < len(customers); i++ {
		resp = append(resp, customers[i].ToDto())
	}

	return resp, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repository.ById(id)
	if err != nil {
		return nil, err
	}

	resp := customer.ToDto()

	return &resp, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
