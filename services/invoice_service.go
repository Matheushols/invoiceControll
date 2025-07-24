package services

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
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

// GenerateInvoicePDF create a PDF with invoice parameters
func GenerateInvoicePDF(invoice models.Invoice) (*gofpdf.Fpdf, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Cabeçalho
	pdf.Cell(40, 10, fmt.Sprintf("Invoice #%d", invoice.Id))
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Social Name: "+invoice.SocialName)
	pdf.Ln(8)
	pdf.Cell(40, 10, "Company Document: "+invoice.CompanyDocument)
	pdf.Ln(8)
	pdf.Cell(40, 10, "Date: "+invoice.Date.Format("02/01/2006"))
	pdf.Ln(8)
	pdf.Cell(40, 10, "Due Date: "+invoice.DueDate.Format("02/01/2006"))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Amount: R$ %.2f", invoice.Amount))
	pdf.Ln(8)
	pdf.Cell(40, 10, "NFS Number: "+invoice.NumNfs)
	pdf.Ln(8)
	pdf.Cell(40, 10, "Bar Code: "+invoice.BarCode)
	pdf.Ln(12)

	// Cabeçalho serviços
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(60, 10, "Service")
	pdf.Cell(40, 10, "Description")
	pdf.Cell(20, 10, "Qty")
	pdf.Cell(30, 10, "Unit Price")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 11)
	for _, s := range invoice.Services {
		startY := pdf.GetY()

		pdf.MultiCell(60, 6, s.Name, "", "L", false)
		h1 := pdf.GetY() - startY

		pdf.SetY(startY)
		pdf.SetX(10 + 60)
		pdf.MultiCell(40, 6, s.Description, "", "L", false)
		h2 := pdf.GetY() - startY

		rowHeight := h1
		if h2 > rowHeight {
			rowHeight = h2
		}

		pdf.SetY(startY)
		pdf.SetX(10 + 60 + 40)
		pdf.CellFormat(20, rowHeight, fmt.Sprintf("%d", s.Quantity), "", 0, "C", false, 0, "")
		pdf.CellFormat(30, rowHeight, fmt.Sprintf("R$ %.2f", s.UnitPrice), "", 0, "R", false, 0, "")
		pdf.Ln(rowHeight)
	}

	return pdf, nil
}
