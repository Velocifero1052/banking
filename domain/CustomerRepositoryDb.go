package domain

import (
	errs "banking/errors"
	"banking/logger"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "banking_db"
)

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var findSql string
	var err error

	customers := make([]Customer, 0)

	if status != "" {
		findSql = `select customer_id, name, date_of_birth, city, zipcode, status from customers where status = $1`
		err = d.client.Select(&customers, findSql)
	} else {
		findSql = `select customer_id, name, date_of_birth, city, zipcode, status from customers`
		err = d.client.Select(&customers, findSql)
	}

	if err != nil {
		logger.Error("error while executing query")
		return nil, errs.NewInternalServerError("error while executing query")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {

	customerSql := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id = $1"
	var customer Customer

	err := d.client.Get(&customer, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errs.NewInternalServerError("unexpected database error")
	}

	return &customer, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	client, err := sqlx.Open("postgres", psqlConn)
	CheckError(err)

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
