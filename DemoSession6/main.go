package main

import (
	"config"
	"entities"
	"fmt"
	"models"
)

func SS7Demo1() {
	db, err := config.Connectss7()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModelss7{
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

func SS7Demo2() {
	db, err := config.Connectss7()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModelss7{
			Db: db,
		}
		product, err2 := productModel.Find(4)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(product.ToString())
			fmt.Println("-------------------")
		}
	}
}

func SS7Demo3() {
	db, err := config.Connectss7()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModelss7{
			Db: db,
		}
		products, err2 := productModel.Search(0, 2000)
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

func SS7Demo4() {
	db, err := config.Connectss7()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModelss7{
			Db: db,
		}
		products, err2 := productModel.FindKeyWord("name")
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

func SS7Demod() {
	db, err := config.Connectss7()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModelss7{
			Db: db,
		}
		product := entities.Productss7{
			Name:     "name_new",
			Price:    12.5,
			Quantity: 1000,
		}
		err2 := productModel.Create(&product)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(product.ToString())
			fmt.Println("-------------------")

		}
	}
}

func SS7Demo8() {
	db, err := config.Connectss7()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModelss7{
			Db: db,
		}

		effect, err2 := productModel.Delete(5)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("Row Affected:", effect)
			fmt.Println("-------------------")

		}
	}
}

func SS7Demo9() {
	db, err := config.Connectss7()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModelss7{
			Db: db,
		}
		product, err2 := productModel.Find(4)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			product.Name = "New Name"
			product.Price = 32
			product.Quantity = 300
			effect, err2 := productModel.Update(product)
			if err2 != nil {
				fmt.Println(err2)
			} else {
				fmt.Println("Row Affected:", effect)
				fmt.Println("-------------------")

			}
		}

	}
}

func Demo1() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Failed")
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			fmt.Println(err)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
			}
		}
	}
}
func Demo2() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Failed")
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product, err2 := productModel.Find(1)
		if err2 != nil {
			fmt.Println(err)
		} else {
			fmt.Println(product.ToString())
		}
	}
}
func Demo3() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Failed")
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product := entities.Product{
			Name:     "tu lanh",
			Price:    111,
			Quantity: 22,
		}
		err2 := productModel.Create(&product)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(product.ToString())
		}
	}
}
func Demo4() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Failed")
	} else {
		productModel := models.ProductModel{
			Db: db,
		}

		rows, err2 := productModel.Delete(2)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(rows)
		}
	}
}

func Demo5() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Failed")
	} else {
		productModel := models.ProductModel{
			Db: db,
		}

		product, err2 := productModel.Find(3)
		fmt.Println(product.ToString())
		if err2 != nil {
			fmt.Println(err2)
		} else {
			product.Name = "abc"
			product.Price = 999
			rowAffected, err3 := productModel.Update(product)
			if err3 != nil {
				fmt.Println(err3)
			} else {
				fmt.Println("Rows Affected", rowAffected)
			}
		}
	}
}
func main() {
	SS7Demo9()
}
