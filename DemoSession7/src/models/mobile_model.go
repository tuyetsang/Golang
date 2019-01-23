package models

import (
	"entities"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MobileModel struct {
	Db *mgo.Database
}

func (mobileModel MobileModel) FindAll() ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{}).All(&mobiles)
	return mobiles, err
}

func (mobileModel MobileModel) Search(key string) ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{
		"name": bson.RegEx{
			Pattern: "" + key,
			Options: "i",
		},
	}).All(&mobiles)
	return mobiles, err
}

func (mobileModel MobileModel) Search3(key string) ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{
		"name": bson.RegEx{
			Pattern: key,
			Options: "i",
		},
	}).All(&mobiles)
	return mobiles, err
}

func (mobileModel MobileModel) Condition1(status bool) ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{"status": status}).All(&mobiles)
	return mobiles, err
}
func (mobileModel MobileModel) Condition2(min float64) ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{
		"price": bson.M{"$gte": min},
	}).All(&mobiles)
	return mobiles, err
}
func (mobileModel MobileModel) Condition3(min, max float64) ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{
		"$and": []bson.M{
			bson.M{"price": bson.M{"$gte": min}},
			bson.M{"price": bson.M{"$lte": max}},
		},
	}).All(&mobiles)
	return mobiles, err
}

func (mobileModel MobileModel) Create(mobile *entities.Mobile) error {
	err := mobileModel.Db.C("mobile").Insert(&mobile)
	return err
}

func (mobileModel MobileModel) Delete(id string) error {
	err := mobileModel.Db.C("mobile").RemoveId(bson.ObjectIdHex(id))
	return err
}

func (mobileModel MobileModel) Update(mobile entities.Mobile) error {
	err := mobileModel.Db.C("mobile").UpdateId(mobile.Id, mobile)
	return err
}

func (mobileModel MobileModel) Find(id string) (entities.Mobile, error) {
	var mobile entities.Mobile
	err := mobileModel.Db.C("mobile").FindId(bson.ObjectIdHex(id)).One(&mobile)
	return mobile, err
}
func (mobileModel MobileModel) Limit1(n int) ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{}).Limit(n).All(&mobiles)
	return mobiles, err
}
func (mobileModel MobileModel) Limit2(n int) ([]entities.Mobile, error) {
	var mobiles []entities.Mobile
	err := mobileModel.Db.C("mobile").Find(bson.M{}).Sort("-price").Limit(n).All(&mobiles)
	return mobiles, err
}
