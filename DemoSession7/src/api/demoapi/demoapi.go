package demoapi

import (
	"fmt"
	"net/http"
)

func Demo1(response http.ResponseWriter, request *http.Request) {
	fmt.Println(response, "Demo1Api")
}
func Demo2(response http.ResponseWriter, request *http.Request) {
	fmt.Println(response, "Demo2Api")
}
