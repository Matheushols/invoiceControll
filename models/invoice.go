package models

import "time"

type Service struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Invoice struct {
	Id              int       `json:"id"`
	SocialName      string    `json:"socialName"`
	CompanyDocument string    `json:"companyDocument"`
	Date            time.Time `json:"date"`
	DueDate         time.Time `json:"dueDate"`
	Amount          float64   `json:"amount"`
	NumNfs          string    `json:"numNfs"`
	BarCode         string    `json:"barCode"`
	Services        []Service `json:"services"`
}
