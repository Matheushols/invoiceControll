package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"invoiceControll/models"
	"invoiceControll/services"
	"io"
	"net/http"
	"strconv"
)

var Invoices = make([]models.Invoice, 0)

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

// ListInvoiceHandler used to list all invoices created
func ListInvoicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Invoices)
}

// GetInvoiceHandler search an invoice by id
func GetInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
		return
	}
	foundInvoice, _ := services.SearchInvoiceById(Invoices, id)
	if foundInvoice.Id == 0 {
		http.Error(w, "Invoice not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(foundInvoice)
}

func UpdateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição.", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var invoice models.Invoice
	err = json.Unmarshal(body, &invoice)
	if err != nil {
		http.Error(w, "Erro ao ler o parse de JSON.", http.StatusBadRequest)
		return
	}
	invoice.Id = id
	foundInvoice, index := services.SearchInvoiceById(Invoices, id)
	if foundInvoice.Id == 0 {
		http.Error(w, "Invoice not found", http.StatusNotFound)
		return
	}
	Invoices[index] = invoice
	json.NewEncoder(w).Encode(invoice)
}

func DeleteInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
		return
	}
	foundInvoice, index := services.SearchInvoiceById(Invoices, id)
	if foundInvoice.Id == 0 {
		http.Error(w, "Invoice not found", http.StatusNotFound)
		return
	}
	Invoices = append(Invoices[:index], Invoices[index+1:]...)
	foundInvoice, index = services.SearchInvoiceById(Invoices, id)
	fmt.Fprintf(w, "Invoice deleted")
}
