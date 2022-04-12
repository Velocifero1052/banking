package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	customerHandlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", customerHandlers.getAllCustomers)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandlers.getCustomer)

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8080", router))

}
