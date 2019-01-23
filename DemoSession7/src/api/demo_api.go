package apis

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Demo1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello World")
}

func Demo2(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	fmt.Fprint(response, "<b><i><u>Hello World</u></i></b>")
}

func Demo3(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fullName := vars["fullName"]
	fmt.Fprint(response, "Hello"+fullName)
}
