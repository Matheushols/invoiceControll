package main

import (
	"fmt"
	"invoiceControll/router"
	"net/http"
)

func main() {
	r := router.Initialize()
	fmt.Println("API is running!")
	http.ListenAndServe(":8080", r)
}
