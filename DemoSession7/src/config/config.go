package config

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

func Connect() (*mgo.Database, error) {
	host := "mongodb://localhost:27017"
	dbName := "demo_session_7"
	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		db := session.DB(dbName)
		return db, nil
	}
}
