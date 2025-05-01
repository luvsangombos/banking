package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/luvsangombos/banking/errs"
	"github.com/luvsangombos/banking/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select * from customers where customer_id = ?"
	var c Customer
	err := d.client.Get(&c, customerSql, id)

	// err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected db error")
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)
	var err error

	if status == "" {
		findAllSql := "select * from customers"
		err = d.client.Select(&customers, findAllSql)

	} else {
		findAllSql := "select * from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected db error")
	}
	// result := make([]Customer, 0)
	// err = sqlx.StructScan(rows, &result)
	// if err != nil {
	// 	logger.Error("Error while quering customer table" + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected db error")
	// }

	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	// 	if err != nil {
	// 		logger.Error("Error while quering customer table" + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unexpected db error")
	// 	}
	// 	result = append(result, c)

	// }
	return customers, nil
}

func NewCustomerRepositoryDB(client *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client}

}
