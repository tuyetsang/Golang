package entities

import (
	"fmt"
)

type Productss7 struct {
	Id       int64
	Name     string
	Price    float64
	Quantity int
}

func (product Productss7) ToString() string {
	return fmt.Sprintf("Id:%d\nName:%s\nPrice:%f\nQuantity:%d", product.Id, product.Name, product.Price, product.Quantity)
}
