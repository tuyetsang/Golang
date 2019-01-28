package demoapi

import (
	"fmt"
	"net/http"
)

func Demo1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Demo 1 API")
}

func Demo2(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Demo 2 API")
}
