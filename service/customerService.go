package service

import (
	"github.com/luvsangombos/banking/domain"
	"github.com/luvsangombos/banking/dto"
	"github.com/luvsangombos/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	ById(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	res, err := s.repo.FindAll(status)

	if err != nil {
		return nil, err
	}

	dtoArray := make([]dto.CustomerResponse, 0)

	for _, val := range res {
		dto := val.ToDto()
		dtoArray = append(dtoArray, dto)
	}

	return dtoArray, nil
}

func (s DefaultCustomerService) ById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
