package entities

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
	Fullname string        `json:"fullname" bson:"fullname"`
}

func (account Account) ToString() string {
	return fmt.Sprintf("Username:%s\nPassword:%s", account.Username, account.Password)
}
