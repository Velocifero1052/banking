package app

import (
	"banking/service"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zipCode" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(writer http.ResponseWriter, request *http.Request) {

	customers, err := ch.service.GetAllCustomer()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(customers)

	contentType := request.Header.Get("Content-Type")
	if contentType == "application/xml" {
		writer.Header().Add("Content-Type", "application/xml")
		err := xml.NewEncoder(writer).Encode(customers)
		if err != nil {
			return
		}
	} else if contentType == "application/json" {
		writer.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode(customers)
		if err != nil {
			return
		}
	}

}

func (ch *CustomerHandlers) getCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		_, err := fmt.Fprintf(writer, err.Error())
		if err != nil {
			return
		}
	} else {
		writer.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode(customer)
		if err != nil {
			return
		}
	}

}
