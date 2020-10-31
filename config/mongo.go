package config

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

// GetMongoDB returns mongo db
func GetMongoDB() (*mgo.Database, error) {
	host := os.Getenv("localhost:27017")
	dbName := os.Getenv("goapi")

	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)
	return db, nil
}
