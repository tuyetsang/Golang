package entities

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Mobile struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Price    float64       `json:"price" bson:"price"`
	Quantity int           `json:"quantity" bson:"quantity"`
}

func (mobile Mobile) ToString() string {
	return fmt.Sprintf("Id:%s\nName:%s\nPrice:%f\nQuantity:%d", mobile.Id, mobile.Name, mobile.Price, mobile.Quantity)
}
