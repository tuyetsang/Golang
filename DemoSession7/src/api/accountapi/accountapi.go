package accountapi

import (
	"encoding/json"
	"entities"
	"net/http"
)

func respondWithJson(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
func respondWithError(response http.ResponseWriter, statusCode int, msg string) {
	respondWithJson(response, statusCode, map[string]string{"error": msg})
}
func Find(response http.ResponseWriter, request *http.Request) {
	account := entities.Account{
		Username: "sang",
		Password: "12345678",
		Age:      24,
		Status:   true,
	}
	respondWithJson(response, http.StatusOK, account)
}
func FindAll(response http.ResponseWriter, request *http.Request) {
	accounts := []entities.Account{
		entities.Account{
			Username: "sang",
			Password: "12345678",
			Age:      24,
			Status:   true,
		},
		entities.Account{
			Username: "tri",
			Password: "123456",
			Age:      25,
			Status:   false,
		},
		entities.Account{
			Username: "huu",
			Password: "1234567",
			Age:      26,
			Status:   true,
		},
	}
	respondWithJson(response, http.StatusOK, accounts)
}
