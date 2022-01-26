package customer

import (
	"encoding/json"
	"net/http"

	"github.com/cheranga/inventoryservice/util"
)

func customerHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		customerId, err := util.GetIdFromUrl(request)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		dto, dtoError := GetCustomerById(customerId)
		if dtoError != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var responseBytes, _ = json.Marshal(dto)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(responseBytes)
		writer.WriteHeader(http.StatusOK)
	}
}
