package app

import (
	"banking/domain"
	"banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	customersDb := domain.NewCustomerRepositoryDb()
	//wiring
	customerHandlers := CustomerHandlers{service.NewCustomerService(customersDb)}

	router.HandleFunc("/customers", customerHandlers.getAllCustomers)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandlers.getCustomer)

	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
