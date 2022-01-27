package customer

import (
	"net/http"

	"github.com/cheranga/inventoryservice/common"
	"github.com/cheranga/inventoryservice/util"
)

func customerHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		customerId, err := util.GetIdFromUrl(request)
		if err != nil {
			RenderGetCustomerByIdResponse(writer, GetCustomerByIdResponse{}, common.ErrorResponse{ErrorCode: common.InvalidCustomerId, ErrorMessage: common.InvalidCustomerIdMessage})
			return
		}

		var appRequest = GetCustomerByIdRequest{
			CustomerId: customerId,
		}

		dto, dtoError := GetCustomer(appRequest)
		RenderGetCustomerByIdResponse(writer, dto, dtoError)

	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func customersHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		id, err := util.GetIdFromUrl(request)
		if err != nil {
			RenderGetCustomersResponse(writer, GetCustomersResponse{}, common.ErrorResponse{ErrorCode: common.InvalidPageId, ErrorMessage: common.InvalidPageIdMessage})
			return
		}

		var appRequest = GetCustomersByPageIdRequest{
			PageId: id,
		}
		dto, dtoError := GetCustomers(appRequest)

		RenderGetCustomersResponse(writer, dto, dtoError)

	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}
