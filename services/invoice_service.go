package services

import (
	"invoiceControll/models"
	"sync"
)

var (
	lastID int
	mutex  sync.Mutex
)

// CreateInvoice generates a new invoice ID and returns the invoice with ID set
func CreateInvoice(data models.Invoice) models.Invoice {
	mutex.Lock()
	lastID++
	data.Id = lastID
	mutex.Unlock()
	return data
}

// SearchInvoiceById used to get a invoice by id
func SearchInvoiceById(Invoices []models.Invoice, id int) (models.Invoice, int) {
	var findedInvoice models.Invoice
	index := -1
	for i, invoice := range Invoices {
		if id == invoice.Id {
			findedInvoice = invoice
			index = i
			break
		}
	}
	return findedInvoice, index
}
