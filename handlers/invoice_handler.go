package handlers

import (
	"encoding/json"
	"invoiceControll/models"
	"invoiceControll/services"
	"io"
	"net/http"
)

var Invoices []models.Invoice

// CreateInvoiceHandler is a handler to create a new invoice
func CreateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição.", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var invoice models.Invoice
	err = json.Unmarshal(body, &invoice)
	if err != nil {
		http.Error(w, "Erro ao fazer o parse de JSON.", http.StatusBadRequest)
		return
	}

	invoice = services.CreateInvoice(invoice)

	Invoices = append(Invoices, invoice)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}
