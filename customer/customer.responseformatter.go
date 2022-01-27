package customer

import (
	"encoding/json"
	"net/http"
)

func RenderCustomerResponse(writer http.ResponseWriter, data GetCustomerByIdResponse, err error) {
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
