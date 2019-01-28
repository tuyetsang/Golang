package models

import (
	"entities"

	"gopkg.in/mgo.v2/bson"

	"golang.org/x/crypto/bcrypt"

	mgo "gopkg.in/mgo.v2"
)

type AccountModel struct {
	Db *mgo.Database
}

func (accountModel AccountModel) Create(account *entities.Account) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hash)
	return accountModel.Db.C("account").Insert(&account)
}

func (accountModel AccountModel) CheckUsernameAndPassword(username, password string) bool {
	var account entities.Account
	err := accountModel.Db.C("account").Find(bson.M{"username": username}).One(&account)
	if err != nil {
		return false
	} else {
		return bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)) == nil
	}
}
