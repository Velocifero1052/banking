package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "banking_db"
)

func Start() {
	router := mux.NewRouter()
	client := connectAndGetClientDbToDb(host, port, user, password, dbname)
	customersDb := domain.NewCustomerRepositoryDb(client)
	//accountsDb := domain.NewAccountRepositoryDb(connection)

	customerHandlers := CustomerHandlers{service.NewCustomerService(customersDb)}

	router.HandleFunc("/customers", customerHandlers.getAllCustomers)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandlers.getCustomer)

	log.Fatal(http.ListenAndServe("localhost:8080", router))

}

func connectAndGetClientDbToDb(host string, port int, user, password, dbname string) *sqlx.DB {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	client, err := sqlx.Open("postgres", psqlConn)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
