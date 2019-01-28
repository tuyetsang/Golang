package main

import (
	apis "api"
	"api/accountapi"
	"api/demo1api"
	"api/demo2api"
	"api/demoapi"
	"api/mobileapi"
	"api/productapi"
	"config"
	"entities"
	"fmt"
	"jwtauth"
	"middlewares/basic2Auth"
	"middlewares/jwtmiddleware"
	"middlewares/log2middleware"
	"middlewares/logmiddleware"
	"models"
	"net/http"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2/bson"
)

func Demo1_1() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("-------------------")
			}
		}
	}
}

func Demo1_2() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product, err2 := productModel.Find("5c3d9e02de2fab661ea4fb33")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(product.ToString())
			fmt.Println("-------------------")
		}
	}
}

func Demo1_3() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product := entities.Product{
			Id:       bson.NewObjectId(),
			Name:     "ABCNew",
			Price:    44,
			Quantity: 2,
			Status:   true,
		}
		err2 := productModel.Create(&product)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(product.ToString())
		}
	}
}

func Demo1_4() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		err2 := productModel.Delete("5c4184eb4912c611fc910745")
		if err2 != nil {
			fmt.Println(err2)
		}
	}
}

func Demo1_5() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product, err2 := productModel.Find("5c3d9e02de2fab661ea4fb33")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			product.Name = "DEF"
			product.Price = 9999
			err3 := productModel.Update(product)
			if err3 != nil {
				fmt.Println(err3)
			} else {
				fmt.Println("Done!")
			}
		}

	}
}

func Demo1_6() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Condition1(false)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("-------------------")
			}
		}
	}
}

func Demo1_7() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Condition2(5000)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("-------------------")
			}
		}
	}
}

func Demo1_8() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Condition3(0, 5000)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("-------------------")
			}
		}
	}
}

func Demo1_9() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Condition4("no")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("-------------------")
			}
		}
	}
}

func Demo1_10() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Sort()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("-------------------")
			}
		}
	}
}

func Demo1_11() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Limit1(3)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("-------------------")
			}
		}
	}
}

func Demo1() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		mobiles, err2 := mobileModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, mobile := range mobiles {
				fmt.Println(mobile.ToString())
				fmt.Println("-------------------")
			}
		}
	}
}

func Demo2() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		mobiles, err2 := mobileModel.Condition1(true)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, mobile := range mobiles {
				fmt.Println(mobile.ToString())
				fmt.Println("-------------------")
			}
		}
	}
}
func Demo3() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		mobiles, err2 := mobileModel.Condition2(10)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, mobile := range mobiles {
				fmt.Println(mobile.ToString())
				fmt.Println("-------------------")
			}
		}
	}
}
func Demo4() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		mobiles, err2 := mobileModel.Search2(4, 20)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, mobile := range mobiles {
				fmt.Println(mobile.ToString())
				fmt.Println("-------------------")
			}
		}
	}
}

func demo5() {
	router := mux.NewRouter()
	router.HandleFunc("/api/demo/demo1", apis.Demo1).Methods("GET")
	router.HandleFunc("/api/demo/demo2", apis.Demo2).Methods("GET")
	router.HandleFunc("/api/demo/demo3/{fullName}", apis.Demo3).Methods("GET")
	router.HandleFunc("/api/mobile/", mobileapi.GetAll).Methods("GET")
	router.HandleFunc("/api/mobile/search1/{keyword}", mobileapi.Search).Methods("GET")
	router.HandleFunc("/api/mobile/search2/{min}/{max}", mobileapi.Search2).Methods("GET")
	router.HandleFunc("/api/mobile/create", mobileapi.Create).Methods("POST")
	router.HandleFunc("/api/mobile/update/{id}", mobileapi.Update).Methods("POST")
	router.HandleFunc("/api/mobile/delete/{id}", mobileapi.Delete).Methods("DELETE")
	router.HandleFunc("/api/product/create", productapi.Create).Methods("POST")
	router.HandleFunc("/api/product/edit", productapi.Update).Methods("PUT")
	router.HandleFunc("/api/product/delete/{id}", productapi.Delete).Methods("DELETE")
	err := http.ListenAndServe(":2803", router)
	if err != nil {
		fmt.Println(err)
	}
}

func Demo7() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err)
	} else {
		mobileModel := models.MobileModel{
			Db: db,
		}
		mobiles, err2 := mobileModel.Limit1(1)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, mobile := range mobiles {
				fmt.Println(mobile.ToString())
				fmt.Println("-------------------")
			}
		}
	}
}

func main() {
	router := mux.NewRouter()
	router.Use(
		logmiddleware.Logmiddleware,
		log2middleware.DateLog,
	)
	router.HandleFunc("/api/demo/demo1", demo1api.Demo1).Methods("GET")
	router.HandleFunc("/api/demo/demo2", demo1api.Demo2).Methods("GET")
	router.HandleFunc("/api/demo/demo3/{fullName}", demo1api.Demo3).Methods("GET")
	router.HandleFunc("/api/demo/sum/{num1}/{num2}", demo1api.Sum).Methods("GET")
	router.HandleFunc("/api/demo/demo4", demo1api.Demo4).Methods("POST")
	router.HandleFunc("/api/demo/demo5", demo1api.Demo5).Methods("PUT")
	router.HandleFunc("/api/demo/demo6/{id}", demo1api.Demo6).Methods("DELETE")
	router.HandleFunc("/api/account/create", accountapi.CreateAccount).Methods("POST")
	router.HandleFunc("/api/jwt/generratetoken", jwtauth.GenerateToken).Methods("POST")
	router.Handle("/api/demo/demo1", jwtmiddleware.JWTAuth(http.HandlerFunc(demoapi.Demo1))).Methods("GET")

	router.Handle("/api/demo/demo2", jwtmiddleware.JWTAuth(http.HandlerFunc(demoapi.Demo1))).Methods("GET")
	// router.HandleFunc("/api/account/find", accountapi.Find).Methods("GET")
	// router.HandleFunc("/api/account/findall", accountapi.FindAll).Methods("GET")
	router.HandleFunc("/api/mobile/findall", mobileapi.FindAll).Methods("GET")
	router.HandleFunc("/api/mobile/find/{id}", mobileapi.Search1).Methods("GET")
	router.HandleFunc("/api/mobile/search1/{keyword}", mobileapi.SearchKey).Methods("GET")
	router.HandleFunc("/api/mobile/search2/{min}/{max}", mobileapi.SearchMinMax).Methods("GET")
	router.HandleFunc("/api/mobile/create", mobileapi.CreateMobile).Methods("POST")
	router.HandleFunc("/api/mobile/update", mobileapi.UpdateMobile).Methods("PUT")
	router.HandleFunc("/api/mobile/delete/{id}", mobileapi.DeleteMobile).Methods("DELETE")

	// router.Handle("/api/demo2/hello", log2middleware.Log2middleware(logmiddleware.Logmiddleware(http.HandlerFunc(demo2api.Hello)))).Methods("GET")
	router.Handle("/api/demo2/hello", basic2Auth.BasicAuth2(http.HandlerFunc(demo2api.Hello))).Methods("GET")
	router.Handle("/api/demo2/hi", logmiddleware.Logmiddleware(http.HandlerFunc(demo2api.Hi))).Methods("GET")
	router.HandleFunc("/api/mobile/register", accountapi.CreateAccount).Methods("POST")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}
