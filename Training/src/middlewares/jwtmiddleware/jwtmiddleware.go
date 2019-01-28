package jwtmiddleware

import (
	"encoding/json"
	"net/http"
	"security"

	jwt "github.com/dgrijalva/jwt-go"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		stringToken := request.Header.Get("token")
		if stringToken == "" {
			respondWithError(response, http.StatusUnauthorized, "Unauthorized")
			return
		} else {
			result, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
				return []byte(security.PrivateKey), nil
			})
			if err != nil {
				respondWithError(response, http.StatusUnauthorized, err.Error())
				return
			} else {
				if result.Valid {
					next.ServeHTTP(response, request)
				} else {
					respondWithError(response, http.StatusUnauthorized, "Invalid Token")
				}
			}
		}
	})
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
