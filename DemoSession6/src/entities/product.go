package entities

import (
	"fmt"
)

type Product struct {
	Id       int64
	Name     string
	Price    float64
	Quantity int64
}

func (product Product) ToString() string {
	return fmt.Sprintf("Id:%d\nName:%s\nPrice:%f\nQuantity:%d", product.Id, product.Name, product.Price, product.Quantity)
}
