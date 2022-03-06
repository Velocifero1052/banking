package app

import (
	"banking/service"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"time"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zipCode" xml:"zipcode"`
}

type DefaultTime struct {
	CurrentTime string `json:"current_time"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(writer http.ResponseWriter, request *http.Request) {
	customer, err := ch.service.GetAllCustomer()

	if err != nil {
		return
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

func addNewCustomer(writer http.ResponseWriter, request *http.Request) {

	_, err := fmt.Fprint(writer, "post request received")

	if err != nil {
		return
	}

}

func getTime(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	timeZones := params["tz"]

	argsCount := len(timeZones)
	writer.Header().Add("Content-Type", "application/json")

	if argsCount == 0 {
		defaultTime := DefaultTime{time.Now().UTC().String()}
		err := json.NewEncoder(writer).Encode(defaultTime)
		if err != nil {
			return
		}
	} else {

		zones := timeZones[0]
		areas := strings.Split(zones, ",")

		answer := map[string]string{}

		for _, area := range areas {
			loc, err := time.LoadLocation(area)

			if err != nil {
				writer.Header().Set("Content-Type", "text/plain")
				writer.WriteHeader(http.StatusNotFound)
				_, err := fmt.Fprint(writer, "invalid timezone")
				if err != nil {
					return
				}
				return
			}

			answer[area] = time.Now().In(loc).String()
		}

		err := json.NewEncoder(writer).Encode(answer)
		if err != nil {
			return
		}

	}
}
