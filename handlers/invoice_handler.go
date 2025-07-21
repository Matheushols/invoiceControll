package handlers

import (
	"fmt"
	"invoiceControll/services"
	"net/http"
)

// HelloInvoiceHandler is a test handler
func HelloInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	message := services.HelloTeste()
	fmt.Fprintln(w, message)
}
