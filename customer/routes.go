package customer

import (
	"net/http"

	"github.com/cheranga/inventoryservice/middleware"
)

func SetupRoutes() {
	customerHandler := http.HandlerFunc(customerHandler)
	customersHandler := http.HandlerFunc(customersHandler)

	http.Handle("/all/customers/", middleware.MiddlewareHandler(customersHandler))
	http.Handle("/customers/", middleware.MiddlewareHandler(customerHandler))
}
