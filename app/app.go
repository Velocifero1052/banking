package app

import (
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/getCustomer", getCustomer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
