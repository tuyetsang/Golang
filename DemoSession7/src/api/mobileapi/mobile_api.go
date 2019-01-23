package mobileapi

import (
	"config"
	"encoding/json"
	"entities"
	"fmt"
	"models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// function name begin with lower word -> private and upper word -> public
func responseWithJson(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJson(response, statusCode, map[string]string{"error": msg})
}

func GetAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		mobiles, err2 := mobileModel.FindAll()
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJson(response, http.StatusOK, mobiles)
		}
	}
}

func Search(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := config.Connect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		mobiles, err2 := mobileModel.Search(keyword)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJson(response, http.StatusOK, mobiles)
		}
	}
}

func Search2(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	min_price := vars["min"]
	max_price := vars["max"]
	min_float, _ := strconv.ParseFloat(min_price, 2)
	max_float, _ := strconv.ParseFloat(max_price, 2)
	db, err := config.Connect()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		mobiles, err2 := mobileModel.Condition3(min_float, max_float)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJson(response, http.StatusOK, mobiles)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	var mobile_param entities.Product
	err := json.NewDecoder(request.Body).Decode(&mobile_param)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		db, err2 := config.Connect()
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			mobileModel := models.MobileModel{
				Db: db,
			}
			mobile := entities.Mobile{
				Id:       bson.NewObjectId(),
				Name:     mobile_param.Name,
				Price:    mobile_param.Price,
				Quantity: mobile_param.Quantity,
			}
			err3 := mobileModel.Create(&mobile)
			if err3 != nil {
				responseWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				responseWithJson(response, http.StatusOK, mobile)
			}
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	mobile_id := vars["id"]
	var mobile_param entities.Mobile
	err := json.NewDecoder(request.Body).Decode(&mobile_param)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		db, err2 := config.Connect()
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			mobileModel := models.MobileModel{
				Db: db,
			}
			mobile, err3 := mobileModel.Find(mobile_id)
			if err3 != nil {
				responseWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				mobile.Name = mobile_param.Name
				mobile.Price = mobile_param.Price
				mobile.Quantity = mobile_param.Quantity
				err4 := mobileModel.Update(mobile)
				if err4 != nil {
					responseWithError(response, http.StatusBadRequest, err4.Error())
				} else {
					responseWithJson(response, http.StatusOK, "Update successfully")
				}
			}

		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		err2 := productModel.Delete(id)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

}
