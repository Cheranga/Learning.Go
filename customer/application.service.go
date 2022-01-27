package customer

import "github.com/cheranga/inventoryservice/common"

var customerHttpService ICustomerHttpService

func GetCustomer(request GetCustomerByIdRequest) (GetCustomerByIdResponse, common.ErrorResponse) {

	customerHttpService = CustomerHttpService{}

	customerResponse, customerError := customerHttpService.GetCustomerById(request.CustomerId)
	return customerResponse, customerError
}

func GetCustomers(request GetCustomersByPageIdRequest) (GetCustomersResponse, common.ErrorResponse) {

	customerHttpService = CustomerHttpService{}

	customerResponse, customerError := customerHttpService.GetAllCustomers(request.PageId)

	if !customerError.IsValid() {
		return customerResponse, customerError
	}

	if len(customerResponse.Data) == 0 {
		customerError.ErrorCode = common.CustomerNotFound
		customerError.ErrorMessage = common.CustomerNotFoundMessage
	}

	return customerResponse, customerError
}
