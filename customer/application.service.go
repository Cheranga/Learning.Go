package customer

var customerHttpService ICustomerHttpService

func GetCustomer(request GetCustomerByIdRequest) (GetCustomerByIdResponse, ErrorResponse) {

	customerHttpService = CustomerHttpService{}

	customerResponse, customerError := customerHttpService.GetCustomerById(request.CustomerId)
	return customerResponse, customerError
}

func GetCustomers(request GetCustomersByPageIdRequest) (GetCustomersResponse, error) {

	customerHttpService = CustomerHttpService{}

	customerResponse, customerError := customerHttpService.GetAllCustomers(request.PageId)
	return customerResponse, customerError
}
