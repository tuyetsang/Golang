package accountapi

import (
	"config"
	"encoding/json"
	"entities"
	"models"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func responseWithJson(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJson(response, statusCode, map[string]string{"error": msg})
}

func CreateAccount(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		accountModel := models.AccountModel{
			Db: db,
		}

		var account entities.Account
		account.Id = bson.NewObjectId()
		err2 := json.NewDecoder(request.Body).Decode(&account)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := accountModel.Create(&account)
			if err3 != nil {
				responseWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				responseWithJson(response, http.StatusOK, account)
			}
		}
	}

}

// func Find(response http.ResponseWriter, request *http.Request) {
// 	account := entities.Account{
// 		Username: "sang",
// 		Password: "12345678",
// 		Age:      24,
// 		Status:   true,
// 	}
// 	respondWithJson(response, http.StatusOK, account)
// }
// func FindAll(response http.ResponseWriter, request *http.Request) {
// 	accounts := []entities.Account{
// 		entities.Account{
// 			Username: "sang",
// 			Password: "12345678",
// 			Age:      24,
// 			Status:   true,
// 		},
// 		entities.Account{
// 			Username: "tri",
// 			Password: "123456",
// 			Age:      25,
// 			Status:   false,
// 		},
// 		entities.Account{
// 			Username: "huu",
// 			Password: "1234567",
// 			Age:      26,
// 			Status:   true,
// 		},
// 	}
// 	respondWithJson(response, http.StatusOK, accounts)
// }
