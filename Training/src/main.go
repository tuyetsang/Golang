package main

import (
	"apis/accountapi"
	"apis/demoapi"
	"apis/jwtauth"
	"fmt"
	"middlewares/jwtmiddleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/account/create", accountapi.Create).Methods("POST")
	router.HandleFunc("/api/jwt/generatetoken", jwtauth.GenerateToken).Methods("POST")

	//router.HandleFunc("/api/demo/demo1", demoapi.Demo1).Methods("GET")
	//router.HandleFunc("/api/demo/demo2", demoapi.Demo2).Methods("GET")

	router.Handle("/api/demo/demo1", jwtmiddleware.JWTAuth(http.HandlerFunc(demoapi.Demo1))).Methods("GET")

	router.Handle("/api/demo/demo2", jwtmiddleware.JWTAuth(http.HandlerFunc(demoapi.Demo1))).Methods("GET")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}

}
