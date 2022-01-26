package customer

const baseUrl = "https://reqres.in/api/users"

var customerHttpService ICustomerHttpService

func GetCustomer(request GetCustomerByIdRequest) (*GetCustomerByIdResponse, error) {

	customerHttpService = CustomerHttpService{}

	customerResponse, customerError := customerHttpService.GetCustomerById(request.CustomerId)
	if customerError != nil {
		return nil, customerError
	}

	return customerResponse, nil
}
