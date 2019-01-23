package entities

import(
	"fmt"
)

type Product struct{
	Id,Name string
	Quantity int
	Price float32
}

func (product Product) ToString() string{
	return fmt.Sprintf("Id: %s\nName:%s\nQuantity:%d\nPrice:%f",product.Id,product.Name,product.Quantity,product.Price)
}

func (product Product) Total() float32{
	return float32(product.Quantity) * float32(product.Price)
}
