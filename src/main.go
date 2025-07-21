package main

import (
	"fmt"
	"invoiceControll/router"
	"net/http"
)

func main() {
	r := router.Initialize()
	http.ListenAndServe(":8080", r)
	fmt.Println("API is running!")
}
