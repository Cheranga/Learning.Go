package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func MiddlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("started processing request")
		start := time.Now()
		handler.ServeHTTP(writer, request)

		fmt.Printf("finished processing %s", time.Since(start))
	})
}
