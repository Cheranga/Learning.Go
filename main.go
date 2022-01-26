package main

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

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

var productList []Product

func init() {
	productsJson := `[
	{
	  "productId": 1,
	  "manufacturer": "Johns-Jenkins",
	  "sku": "p5z343vdS",
	  "upc": "939581000000",
	  "pricePerUnit": "497.45",
	  "quantityOnHand": 9703,
	  "productName": "sticky note"
	},
	{
	  "productId": 2,
	  "manufacturer": "Hessel, Schimmel and Feeney",
	  "sku": "i7v300kmx",
	  "upc": "740979000000",
	  "pricePerUnit": "282.29",
	  "quantityOnHand": 9217,
	  "productName": "leg warmers"
	},
	{ 
	  "productId": 3,
	  "manufacturer": "Swaniawski, Bartoletti and Bruen",
	  "sku": "q0L657ys7",
	  "upc": "111730000000",
	  "pricePerUnit": "436.26",
	  "quantityOnHand": 5905,
	  "productName": "lamp shade"
	}
  ]`

	error := json.Unmarshal([]byte(productsJson), &productList)
	if error != nil {
		log.Fatal(error)
	}
}

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("started processing request")
		start := time.Now()
		handler.ServeHTTP(writer, request)

		fmt.Printf("finished processing %s", time.Since(start))
	})
}

func findProductById(productId int) (*Product, int) {
	for i, product := range productList {
		if product.ProductID == productId {
			return &product, i
		}
	}

	return nil, -1
}

func productHandler(writer http.ResponseWriter, request *http.Request) {
	queryParamters := strings.Split(request.URL.Path, "products/")
	productId, error := strconv.Atoi(queryParamters[len(queryParamters)-1])
	if error != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	product, _ := findProductById(productId)
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
		data, error := json.Marshal(productList)
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

		newProductId := len(productList) + 1
		newProduct.ProductID = newProductId
		productList = append(productList, newProduct)

		writer.WriteHeader(http.StatusCreated)

	}
}

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

	productHandler := http.HandlerFunc(productHandler)
	productListHandler := http.HandlerFunc(productsHandler)

	http.Handle("/products/", middlewareHandler(productHandler))
	http.Handle("/products", middlewareHandler(productListHandler))
	// Commenting so that the handlers can work with the middleware handler
	// http.HandleFunc("/products", middlewareHandler(productListHandler))
	// http.HandleFunc("/products/", productHandler)

	http.Handle("/test", &testHandler{Message: "Hi from handler function"})
	http.HandleFunc("/test2", anotherTestHandler)
	http.ListenAndServe(":5000", nil)
}
