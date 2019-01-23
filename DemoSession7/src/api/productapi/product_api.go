package productapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"parameters"

	"github.com/gorilla/mux"
)

func Create(response http.ResponseWriter, request *http.Request) {
	var product parameters.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Product Information - Create")
		fmt.Println(product.ToString())
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	var product parameters.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Product Information - Update")
		fmt.Println(product.ToString())
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	fmt.Println("id:" + id)
}
