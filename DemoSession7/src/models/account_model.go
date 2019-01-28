package models

import (
	"entities"

	"golang.org/x/crypto/bcrypt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AccountModel struct {
	Db *mgo.Database
}

func (accountModel AccountModel) Create(account *entities.Account) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hash)
	err := accountModel.Db.C("account").Insert(&account)
	return err
}

func (accountModel AccountModel) CheckAuthentication(username, password string) (entities.Account, error) {
	var account entities.Account
	err := accountModel.Db.C("account").Find(bson.M{
		"username": bson.RegEx{
			Pattern: username,
			Options: "i",
		},
	}).One(&account)

	return account, err
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
