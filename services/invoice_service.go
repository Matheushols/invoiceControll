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
