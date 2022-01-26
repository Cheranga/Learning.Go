package customer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const baseUrl = "https://reqres.in/api/users"

func GetCustomerById(customerId int) (*GetCustomerByIdResponse, error) {

	url := fmt.Sprintf("%s/%s", baseUrl, strconv.Itoa(customerId))
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	responseData, responseError := ioutil.ReadAll(response.Body)
	if responseError != nil {
		return nil, responseError
	}

	var customerResponse Customer
	responseDataError := json.Unmarshal(responseData, &customerResponse)
	if responseDataError != nil {
		return nil, responseDataError
	}

	var dto = GetCustomerByIdResponse{
		Id:    customerResponse.Data.Id,
		Email: customerResponse.Data.Email,
	}

	return &dto, nil
}
