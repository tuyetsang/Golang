package basicAuth

import (
	"fmt"
	"net/http"
)

func BasicAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		username, password, ok := request.BasicAuth()
		fmt.Println("username:", username)
		fmt.Println("password:", password)
		fmt.Println("ok:", ok)
		if !ok || !checkUsernameAndPassword(username, password) {
			response.WriteHeader(http.StatusUnauthorized)
			response.Write([]byte("Unauthorized"))
			return
		} else {
			handler.ServeHTTP(response, request)
		}
	})
}
func checkUsernameAndPassword(username, password string) bool {
	return username == "admin" && password == "12345678x@X"
}
