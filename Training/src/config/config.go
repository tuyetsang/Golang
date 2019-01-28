package config

import (
	mgo "gopkg.in/mgo.v2"
)

func Connect() (*mgo.Database, error) {
	host := "mongodb://localhost:27017"
	dbName := "demo_session_9"
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	} else {
		db := session.DB(dbName)
		return db, nil
	}
}
