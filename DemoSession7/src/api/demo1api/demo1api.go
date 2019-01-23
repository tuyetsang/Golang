package demo1api

import (
	"encoding/json"
	"entities"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Demo1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "HelloWorld")
}
func Demo2(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Contend-Type", "text/html")
	fmt.Fprint(response, "<b><i><u>HelloWorld</u></i></b>")
}
func Demo3(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fullName := vars["fullName"]
	fmt.Fprint(response, "Hello\t"+fullName)
}
func Sum(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	num1 := vars["num1"]
	num2 := vars["num2"]
	number1, _ := strconv.ParseInt(num1, 10, 64)
	number2, _ := strconv.ParseInt(num2, 10, 64)
	sum := number1 + number2
	fmt.Fprint(response, sum)
}
func Demo4(response http.ResponseWriter, request *http.Request) {
	var account entities.Account
	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Account-Create")
		fmt.Println(account.ToString())
	}
}
func Demo5(response http.ResponseWriter, request *http.Request) {
	var account entities.Account
	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Account-Update")
		fmt.Println(account.ToString())
	}
}
func Demo6(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	fmt.Println("Id:" + id)
}
