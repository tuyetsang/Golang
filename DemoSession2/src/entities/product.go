package entities

import (
	"fmt"
)

type Product struct {
	Id       string
	Name     string
	Price    float64
	Quantity int
}

func (product Product) ToString() string {
	return fmt.Sprintf("id:%s\nname:%s\nprice:%f\nquantity:%d", product.Id, product.Name, product.Price, product.Quantity)
}
func (product Product) Total() float64 {
	return product.Price * float64(product.Quantity)
}
func NewProduct(id, name string, price float64, quantity int) Product {
	return Product{id, name, price, quantity}

}
