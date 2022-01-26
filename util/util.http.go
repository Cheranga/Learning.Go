package util

import (
	"net/http"
	"strconv"
	"strings"
)

func GetIdFromUrl(request *http.Request) (int, error) {
	queryParamters := strings.Split(request.URL.Path, "customers/")
	id, err := strconv.Atoi(queryParamters[len(queryParamters)-1])

	return id, err
}
