package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()
	/*
		//wiring
		ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

		router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

		router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
		router.HandleFunc("/customer", addNewCustomer).Methods(http.MethodPost)
		router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)*/

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
