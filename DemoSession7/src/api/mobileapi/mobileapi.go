package mobileapi

import (
	"config"
	"encoding/json"
	"entities"
	"models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
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

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		product, err2 := mobileModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}

func Search1(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		product, err2 := mobileModel.Search1("name1")
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}
func Find(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		vars := mux.Vars(request)
		id := vars["id"]
		product, err2 := mobileModel.Find(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}

func SearchKey(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		vars := mux.Vars(request)
		keyword := vars["keyword"]
		product, err2 := mobileModel.Search1(keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}
func SearchMinMax(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		vars := mux.Vars(request)
		min_price := vars["min"]
		max_price := vars["max"]
		min_float, _ := strconv.ParseFloat(min_price, 2)
		max_float, _ := strconv.ParseFloat(max_price, 2)
		product, err2 := mobileModel.Search2(min_float, max_float)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}
func CreateMobile(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}

		var mobile entities.Mobile
		mobile.Id = bson.NewObjectId()
		err2 := json.NewDecoder(request.Body).Decode(&mobile)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := mobileModel.Create(&mobile)
			if err3 != nil {
				responseWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				responseWithJson(response, http.StatusOK, mobile)
			}
		}
	}

}
func UpdateMobile(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}

		var mobile entities.Mobile
		err2 := json.NewDecoder(request.Body).Decode(&mobile)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := mobileModel.Update(&mobile)
			if err3 != nil {
				responseWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				responseWithJson(response, http.StatusOK, mobile)
			}
		}
	}

}
func DeleteMobile(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		vars := mux.Vars(request)
		id := vars["id"]
		err2 := mobileModel.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		}
	}
}
