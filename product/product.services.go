package product

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func productHandler(writer http.ResponseWriter, request *http.Request) {
	productId, err := getProductIdFromUrl(request)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
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
