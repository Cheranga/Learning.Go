package customer

var customerHttpService ICustomerHttpService

func GetCustomer(request GetCustomerByIdRequest) (*GetCustomerByIdResponse, error) {

	customerHttpService = CustomerHttpService{}

	customerResponse, customerError := customerHttpService.GetCustomerById(request.CustomerId)
	if customerError != nil {
		return nil, customerError
	}

	return customerResponse, nil
}

func GetCustomers(request GetCustomersByPageIdRequest) (*GetCustomersResponse, error) {

	customerHttpService = CustomerHttpService{}

	customerResponse, customerError := customerHttpService.GetAllCustomers(request.PageId)
	if customerError != nil {
		return nil, customerError
	}

	return customerResponse, nil
}
