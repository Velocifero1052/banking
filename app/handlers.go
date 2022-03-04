package app

import (
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

func greet(writer http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(writer, "Hello guys!")
	if err != nil {
		return
	}
}

func getAllCustomers(writer http.ResponseWriter, request *http.Request) {
	customer := []Customer{
		{"SomeName", "Tashkent", "100008"},
		{"SomeName", "Tashkent", "100008"},
	}

	contentType := request.Header.Get("Content-Type")
	if contentType == "application/xml" {
		writer.Header().Add("Content-Type", "application/xml")
		err := xml.NewEncoder(writer).Encode(customer)
		if err != nil {
			return
		}
	} else if contentType == "application/json" {
		writer.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode(customer)
		if err != nil {
			return
		}
	}

}

func getCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	_, err := fmt.Fprint(writer, vars["customer_id"])
	if err != nil {
		return
	}
}
