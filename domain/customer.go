package domain

import (
	"github.com/luvsangombos/banking/dto"
	"github.com/luvsangombos/banking/errs"
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
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(id string) (*Customer, *errs.AppError)
}

func (c Customer) StatusAsText() string {
	statusAsTest := "active"
	if c.Status == "1" {
		statusAsTest = "inActive"

	}

	return statusAsTest
}

func (c Customer) ToDto() dto.CustomerResponse {
	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.StatusAsText(),
	}
	return response
}
