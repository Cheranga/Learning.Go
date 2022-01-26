package main

import (
	"net/http"

	"github.com/cheranga/inventoryservice/customer"
	"github.com/cheranga/inventoryservice/product"
)

type testHandler struct {
	Message string
}

func (data *testHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(data.Message))
}

func anotherTestHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hi from HTTP Handler function!"))
}

func main() {

	product.SetupRoutes()
	customer.SetupRoutes()

	// productHandler := http.HandlerFunc(product.ProductHandler)
	// productListHandler := http.HandlerFunc(product.ProductsHandler)

	// http.Handle("/products/", middlewareHandler(productHandler))
	// http.Handle("/products", middlewareHandler(productListHandler))
	// Commenting so that the handlers can work with the middleware handler
	// http.HandleFunc("/products", middlewareHandler(productListHandler))
	// http.HandleFunc("/products/", productHandler)

	http.Handle("/test", &testHandler{Message: "Hi from handler function"})
	http.HandleFunc("/test2", anotherTestHandler)
	http.ListenAndServe(":5000", nil)
}
