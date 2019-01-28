package log2middleware

import (
	"fmt"
	"net/http"
	"time"
)

func DateLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		today := time.Now()
		fmt.Println("date: ", today.Format("01/02/2006 15:04:05"))
		handler.ServeHTTP(response, request)
	})
}
