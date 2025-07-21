package router

import (
	"github.com/gorilla/mux"
	"invoiceControll/handlers"
)

func Initialize() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/invoice/hello", handlers.HelloInvoiceHandler).Methods("GET")

	return r
}
