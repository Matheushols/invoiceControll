package router

import (
	"github.com/gorilla/mux"
	"invoiceControll/handlers"
)

func Initialize() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/invoice", handlers.CreateInvoiceHandler).Methods("POST")
	r.HandleFunc("/invoices", handlers.ListInvoicesHandler).Methods("GET")
	r.HandleFunc("/invoice/{id}", handlers.GetInvoiceHandler).Methods("GET")
	r.HandleFunc("/invoice/{id}", handlers.UpdateInvoiceHandler).Methods("PUT")
	r.HandleFunc("/invoice/{id}", handlers.DeleteInvoiceHandler).Methods("DELETE")

	return r
}
