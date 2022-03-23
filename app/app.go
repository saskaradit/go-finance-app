package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
