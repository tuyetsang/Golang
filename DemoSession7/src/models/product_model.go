package models

import (
	"entities"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductModel struct {
	Db *mgo.Database
}

func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	err := productModel.Db.C("product").Find(bson.M{}).All(&products)
	return products, err
}

func (productModel ProductModel) Find(id string) (entities.Product, error) {
	var product entities.Product
	err := productModel.Db.C("product").FindId(bson.ObjectIdHex(id)).One(&product)
	return product, err
}

func (productModel ProductModel) Create(product *entities.Product) error {
	err := productModel.Db.C("product").Insert(&product)
	return err
}

func (productModel ProductModel) Delete(id string) error {
	err := productModel.Db.C("product").RemoveId(bson.ObjectIdHex(id))
	return err
}

func (productModel ProductModel) Update(product entities.Product) error {
	err := productModel.Db.C("product").UpdateId(product.Id, product)
	return err
}

func (productModel ProductModel) Condition1(status bool) ([]entities.Product, error) {
	var products []entities.Product
	err := productModel.Db.C("product").Find(bson.M{"status": status}).All(&products)
	return products, err
}

func (productModel ProductModel) Condition2(min float64) ([]entities.Product, error) {
	var products []entities.Product
	err := productModel.Db.C("product").Find(bson.M{
		"price": bson.M{"$gt": min},
	}).All(&products)
	return products, err
}

func (productModel ProductModel) Condition3(min, max float64) ([]entities.Product, error) {
	var products []entities.Product
	err := productModel.Db.C("product").Find(bson.M{
		"$and": []bson.M{
			bson.M{"price": bson.M{"$gte": min}},
			bson.M{"price": bson.M{"$lte": max}},
		},
	}).All(&products)
	return products, err
}

func (productModel ProductModel) Condition4(key string) ([]entities.Product, error) {
	var products []entities.Product
	err := productModel.Db.C("product").Find(bson.M{
		"name": bson.RegEx{
			Pattern: "^" + key,
			Options: "i",
		},
	}).All(&products)
	return products, err
}

func (productModel ProductModel) Sort() ([]entities.Product, error) {
	var products []entities.Product
	err := productModel.Db.C("product").Find(bson.M{}).Sort("-price").All(&products)
	return products, err
}

func (productModel ProductModel) Limit1(n int) ([]entities.Product, error) {
	var products []entities.Product
	err := productModel.Db.C("product").Find(bson.M{}).Limit(n).All(&products)
	return products, err
}

func (productModel ProductModel) Limit2(n int) ([]entities.Product, error) {
	var products []entities.Product
	err := productModel.Db.C("product").Find(bson.M{}).Limit(n).Sort("-price").All(&products)
	return products, err
}
