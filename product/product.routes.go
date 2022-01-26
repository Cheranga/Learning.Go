package product

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/cheranga/inventoryservice/middleware"
)

func SetupRoutes() {
	productHandler := http.HandlerFunc(productHandler)
	productListHandler := http.HandlerFunc(productsHandler)

	http.Handle("/products/", middleware.MiddlewareHandler(productHandler))
	http.Handle("/products", middleware.MiddlewareHandler(productListHandler))
}

func getProductIdFromUrl(request *http.Request) (int, error) {
	queryParamters := strings.Split(request.URL.Path, "products/")
	productId, err := strconv.Atoi(queryParamters[len(queryParamters)-1])

	return productId, err
}
