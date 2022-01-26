package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("started processing request")
		start := time.Now()
		handler.ServeHTTP(writer, request)

		fmt.Printf("finished processing %s", time.Since(start))
	})
}

func SetupRoutes() {
	productHandler := http.HandlerFunc(productHandler)
	productListHandler := http.HandlerFunc(productsHandler)

	http.Handle("/products/", middlewareHandler(productHandler))
	http.Handle("/products", middlewareHandler(productListHandler))
}

func productHandler(writer http.ResponseWriter, request *http.Request) {
	queryParamters := strings.Split(request.URL.Path, "products/")
	productId, error := strconv.Atoi(queryParamters[len(queryParamters)-1])
	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	product, _ := GetProductById(productId)
	if product == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	switch request.Method {
	case http.MethodGet:
		productJson, error := json.Marshal(product)
		if error != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(productJson)
		writer.WriteHeader(http.StatusOK)
	case http.MethodPut:
		// TODO
		writer.WriteHeader(http.StatusOK)
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func productsHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		allProducts := GetAllProducts()
		data, error := json.Marshal(allProducts)
		if error != nil {
			log.Fatal(error)
			writer.WriteHeader(http.StatusInternalServerError)
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(data)

	case http.MethodPost:
		var newProduct Product
		requestData, error := ioutil.ReadAll(request.Body)
		if error != nil {
			writer.WriteHeader(http.StatusBadRequest)
		}

		error = json.Unmarshal(requestData, &newProduct)
		if error != nil {
			writer.WriteHeader(http.StatusBadRequest)
		}

		if newProduct.ProductID != 0 {
			writer.WriteHeader(http.StatusBadRequest)
		}

		index := CreateNewProduct(&newProduct)
		if index <= 0 {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusCreated)

	}
}
