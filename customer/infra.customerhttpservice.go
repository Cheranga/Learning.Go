package customer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cheranga/inventoryservice/common"
)

const baseUrl = "https://reqres.in/api/users"

type ICustomerHttpService interface {
	GetCustomerById(customerId int) (GetCustomerByIdResponse, common.ErrorResponse)
	GetAllCustomers(page int) (GetCustomersResponse, common.ErrorResponse)
}

type CustomerHttpService struct {
}

func (customerService CustomerHttpService) GetCustomerById(customerId int) (GetCustomerByIdResponse, common.ErrorResponse) {
	url := fmt.Sprintf("%s/%s", baseUrl, strconv.Itoa(customerId))
	response, err := http.Get(url)
	if err != nil {
		return GetCustomerByIdResponse{}, common.ErrorResponse{ErrorCode: common.CannotConnectToApi, ErrorMessage: common.CannotConnectToApiMessage}
	}

	responseData, _ := ioutil.ReadAll(response.Body)

	var customerResponse Customer
	responseDataError := json.Unmarshal(responseData, &customerResponse)
	if responseDataError != nil {
		return GetCustomerByIdResponse{}, common.ErrorResponse{ErrorCode: common.InvalidResponse, ErrorMessage: common.InvalidResponseMessage}
	}

	if customerId != customerResponse.Data.Id {
		return GetCustomerByIdResponse{}, common.ErrorResponse{ErrorCode: common.CustomerNotFound, ErrorMessage: common.CustomerNotFoundMessage}
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

	return dto, common.ErrorResponse{}
}

func (customerService CustomerHttpService) GetAllCustomers(page int) (GetCustomersResponse, common.ErrorResponse) {
	url := fmt.Sprintf("%s?page=%s", baseUrl, strconv.Itoa(page))
	response, err := http.Get(url)
	if err != nil {
		return GetCustomersResponse{}, common.ErrorResponse{ErrorCode: common.CannotConnectToApi, ErrorMessage: common.CannotConnectToApiMessage}
	}

	responseData, _ := ioutil.ReadAll(response.Body)

	var dto GetCustomersResponse
	responseDataError := json.Unmarshal(responseData, &dto)
	if responseDataError != nil {
		return GetCustomersResponse{}, common.ErrorResponse{ErrorCode: common.InvalidResponse, ErrorMessage: common.InvalidResponseMessage}
	}

	return dto, common.ErrorResponse{}
}
