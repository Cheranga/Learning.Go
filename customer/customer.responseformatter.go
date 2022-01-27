package customer

import (
	"net/http"

	"github.com/cheranga/inventoryservice/common"
)

func RenderGetCustomerByIdResponse(writer http.ResponseWriter, data GetCustomerByIdResponse, err common.ErrorResponse) {

	statusCode := http.StatusOK
	if !err.IsValid() {
		switch err.ErrorCode {
		case common.CustomerNotFound:
			statusCode = http.StatusNotFound
		case common.InvalidCustomerId:
			statusCode = http.StatusBadRequest
		default:
			statusCode = http.StatusInternalServerError
		}
	}

	common.RenderResponse(writer, data, err, statusCode)
}

func RenderGetCustomersResponse(writer http.ResponseWriter, data GetCustomersResponse, err common.ErrorResponse) {
	statusCode := http.StatusOK
	if !err.IsValid() {
		switch err.ErrorCode {
		case common.CustomerNotFound:
			statusCode = http.StatusNotFound
		case common.InvalidPageId:
			statusCode = http.StatusBadRequest
		default:
			statusCode = http.StatusInternalServerError
		}
	}

	common.RenderResponse(writer, data, err, statusCode)
}
