package accountapi

import (
	"config"
	"encoding/json"
	"entities"
	"models"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func Create(response http.ResponseWriter, request *http.Request) {
	var account entities.Account
	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		db, err2 := config.Connect()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			accountModel := models.AccountModel{
				Db: db,
			}
			account.Id = bson.NewObjectId()
			err3 := accountModel.Create(&account)
			if err3 != nil {
				respondWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				respondWithJSON(response, http.StatusOK, account)
			}
		}
	}
}

func respondWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func respondWithError(response http.ResponseWriter, statusCode int, msg string) {
	respondWithJSON(response, statusCode, map[string]string{"error": msg})
}
