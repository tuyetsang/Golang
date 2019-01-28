package jwtauth

import (
	"config"
	"encoding/json"
	"entities"
	"models"
	"net/http"
	"security"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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

func GenerateToken(response http.ResponseWriter, request *http.Request) {
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
			isValid := accountModel.CheckUsernameAndPassword(account.Username, account.Password)
			if !isValid {
				respondWithError(response, http.StatusBadRequest, "Invalid")
			} else {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"username": account.Username,
					"password": account.Password,
					"exp":      time.Now().Add(72 * time.Hour).Unix(),
				})
				tokenString, err2 := token.SignedString([]byte(security.PrivateKey))
				if err2 != nil {
					respondWithError(response, http.StatusBadRequest, err2.Error())
				} else {
					respondWithJson(response, http.StatusOK, entities.JWToken{Token: tokenString})
				}

			}
		}
	}
}
