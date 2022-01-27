package common

import (
	"encoding/json"
	"net/http"
)

func RenderResponse(writer http.ResponseWriter, data interface{}, err ErrorResponse, httpStatusCode int) {

	writer.Header().Set("Content-Type", "application/json")

	dto := DtoResponse{}
	if err.IsValid() {
		dto.Data = data
		responseBytes, _ := json.Marshal(dto)
		writer.WriteHeader(httpStatusCode)
		writer.Write(responseBytes)
		return
	}

	dto.Error = err
	errorData, _ := json.Marshal(dto)

	writer.WriteHeader(httpStatusCode)
	writer.Write(errorData)

}
