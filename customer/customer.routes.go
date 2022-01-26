package customer

import (
	"net/http"

	"github.com/cheranga/inventoryservice/middleware"
)

func SetupRoutes() {
	customerHandler := http.HandlerFunc(customerHandler)

	http.Handle("/customers/", middleware.MiddlewareHandler(customerHandler))
}
