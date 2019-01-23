package entities

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Price    float64       `bson:"price"`
	Quantity int           `bson:"quantity"`
	Status   bool          `bson:"status"`
}

func (product Product) ToString() string {
	return fmt.Sprintf("Id:%s\nName:%s\nPrice:%f\nQuantity:%d\nStatus:%v", product.Id.Hex(), product.Name, product.Price, product.Quantity, product.Status)
}
