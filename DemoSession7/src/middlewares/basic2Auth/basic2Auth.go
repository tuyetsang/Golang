package basic2Auth

import (
	"config"
	"fmt"
	"models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func BasicAuth2(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		username, password, ok := request.BasicAuth()
		fmt.Println("username:", username)
		fmt.Println("password:", password)
		fmt.Println("ok:", ok)
		Passed := checkUser(username, password)
		if !Passed {
			response.WriteHeader(http.StatusUnauthorized)
			response.Write([]byte("Unauthorized"))
			return
		} else {
			handler.ServeHTTP(response, request)
		}
	})
}

func checkUser(_username, _password string) bool {
	db, err := config.Connect()
	if err != nil {
		return false
	} else {
		accountModel := models.AccountModel{
			Db: db,
		}
		account, err2 := accountModel.CheckAuthentication(_username, _password)
		if err2 != nil {
			return false
		} else {
			fmt.Println(account.ToString())
			fmt.Println(bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(_password)) == nil)
			return bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(_password)) == nil
		}
	}
}
