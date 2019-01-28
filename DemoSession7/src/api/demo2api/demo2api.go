package demo2api

import (
	"fmt"
	"net/http"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "HelloWorld")
}
func Hi(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hi")
}
