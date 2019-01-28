package logmiddleware

import (
	"fmt"
	"net/http"
)

func Logmiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		fmt.Println("url: ", request.URL)
		handler.ServeHTTP(response, request)
	})
}
