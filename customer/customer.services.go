package customer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func getIdFromUrl(request *http.Request) (int, error) {
	queryParamters := strings.Split(request.URL.Path, "customers/")
	id, err := strconv.Atoi(queryParamters[len(queryParamters)-1])

	return id, err
}

func customerHandler(writer http.ResponseWriter, request *http.Request) {
	customerId, err := getIdFromUrl(request)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch request.Method {
	case http.MethodGet:

		url := fmt.Sprintf("https://reqres.in/api/users/%s", strconv.Itoa(customerId))
		response, err := http.Get(url)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseData, _ := ioutil.ReadAll(response.Body)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var customerResponse Customer
		responseDataError := json.Unmarshal(responseData, &customerResponse)
		if responseDataError != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var responseBytes, _ = json.Marshal(customerResponse.Data)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(responseBytes)
		writer.WriteHeader(http.StatusOK)
	}
}
