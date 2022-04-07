package domain

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "banking_db"
)

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	db := NewCustomerRepositoryDb()
	findAllSql := `select customer_id, name, date_of_birth, city, zipcode, status from customers`

	rows, err := db.client.Query(findAllSql)

	CheckError(err)
	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer

		err = rows.Scan(&customer.Id, &customer.Name, &customer.DateOfBirth, &customer.City, &customer.ZipCode,
			&customer.Status)

		CheckError(err)

		customers = append(customers, customer)

	}
	return customers, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	client, err := sql.Open("postgres", psqlConn)
	CheckError(err)

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
