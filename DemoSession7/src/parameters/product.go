package parameters

import "fmt"

type Product struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Status   bool    `json:"status"`
}

func (product Product) ToString() string {
	return fmt.Sprintf("Id:%s\nName:%s\nPrice:%f\nQuantity:%d\nStatus:%v", product.Id, product.Name, product.Price, product.Quantity, product.Status)
}
