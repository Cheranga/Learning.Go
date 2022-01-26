package main

import (
	"net/http"

	"github.com/cheranga/inventoryservice/customer"
)

func main() {
	customer.SetupRoutes()
	http.ListenAndServe(":5000", nil)
}
