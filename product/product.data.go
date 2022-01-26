package product

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var productList []Product

func init() {
	file, _ := ioutil.ReadFile("products.json")
	productList = make([]Product, 0)

	error := json.Unmarshal([]byte(file), &productList)
	if error != nil {
		log.Fatal(error)
	}
}

func CreateNewProduct(product *Product) int {
	if product == nil {
		return -1
	}

	var newId = len(productList) + 1
	product.ProductID = newId

	productList = append(productList, *product)
	return newId
}

func GetAllProducts() []Product {
	return productList
}

func GetProductById(productId int) (*Product, int) {
	for i, product := range productList {
		if product.ProductID == productId {
			return &product, i
		}
	}

	return nil, -1
}
