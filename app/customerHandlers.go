package app

import (
	"banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(writer http.ResponseWriter, request *http.Request) {
	vars := request.URL.Query()
	status, containsKey := vars["status"]
	var status_ string

	if !containsKey {
		status_ = ""
	} else {
		status_ = status[0]
	}

	customers, err := ch.service.GetAllCustomer(status_)

	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	}
	writeResponse(writer, http.StatusOK, customers)
}

func (ch *CustomerHandlers) getCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	} else {
		writeResponse(writer, http.StatusOK, customer)
	}

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
