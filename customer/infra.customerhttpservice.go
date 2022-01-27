package customer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const baseUrl = "https://reqres.in/api/users"

type ICustomerHttpService interface {
	GetCustomerById(customerId int) (GetCustomerByIdResponse, error)
	GetAllCustomers(page int) (GetCustomersResponse, error)
}

type CustomerHttpService struct {
}

func (customerService CustomerHttpService) GetCustomerById(customerId int) (GetCustomerByIdResponse, error) {
	url := fmt.Sprintf("%s/%s", baseUrl, strconv.Itoa(customerId))
	response, err := http.Get(url)
	if err != nil {
		return GetCustomerByIdResponse{}, err
	}

	responseData, responseError := ioutil.ReadAll(response.Body)
	if responseError != nil {
		return GetCustomerByIdResponse{}, responseError
	}

	responseContent := string(responseData)
	if responseContent == "{}" {
		return GetCustomerByIdResponse{}, errors.New("No HTTP response")
	}

	var customerResponse Customer
	responseDataError := json.Unmarshal(responseData, &customerResponse)
	if responseDataError != nil {
		return GetCustomerByIdResponse{}, responseDataError
	}

	var dto = GetCustomerByIdResponse{
		Id:        customerResponse.Data.Id,
		Email:     customerResponse.Data.Email,
		FirstName: customerResponse.Data.FirstName,
		LastName:  customerResponse.Data.LastName,
		Avatar:    customerResponse.Data.Avatar,
		Url:       customerResponse.Support.Url,
		Text:      customerResponse.Support.Text,
	}

	return dto, nil
}

func (customerService CustomerHttpService) GetAllCustomers(page int) (GetCustomersResponse, error) {
	url := fmt.Sprintf("%s?page=%s", baseUrl, strconv.Itoa(page))
	response, err := http.Get(url)
	if err != nil {
		return GetCustomersResponse{}, err
	}

	responseData, responseError := ioutil.ReadAll(response.Body)
	if responseError != nil {
		return GetCustomersResponse{}, responseError
	}

	var dto GetCustomersResponse
	responseDataError := json.Unmarshal(responseData, &dto)
	if responseDataError != nil {
		return GetCustomersResponse{}, responseDataError
	}

	return dto, nil
}
