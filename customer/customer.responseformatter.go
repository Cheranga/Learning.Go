package customer

import (
	"encoding/json"
	"net/http"
)

func RenderCustomerResponse(writer http.ResponseWriter, data GetCustomerByIdResponse, err ErrorResponse) {

	writer.Header().Set("Content-Type", "application/json")

	if err.IsValid() {
		responseBytes, _ := json.Marshal(data)
		writer.Write(responseBytes)
		writer.WriteHeader(http.StatusOK)
		return
	}

	switch err.ErrorCode {
	case CustomerNotFound:
		errorData, _ := json.Marshal(err)
		writer.Write(errorData)
		writer.WriteHeader(http.StatusNotFound)
	default:
		errorData, _ := json.Marshal(err)
		writer.Write(errorData)
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func RenderCustomersResponse(writer http.ResponseWriter, data GetCustomersResponse, err error) {
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	var responseBytes, _ = json.Marshal(data)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseBytes)
	writer.WriteHeader(http.StatusOK)
}
