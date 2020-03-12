package models

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

//Database session
type Database struct {
	db *mgo.Database
}

//Connect Returns a database session
func Connect() Database {
	session, err := mgo.Dial("mongo:27017")

	if err != nil {
		fmt.Println("Failed to stablish connection:", err)
	}

	return Database{session.DB("yawoen")}
}

//Close current session
func (d Database) Close() {
	d.db.Session.Close()
}

//Collection Access collection object of a database
func (d Database) Collection(collection string) *mgo.Collection {
	return d.db.C(collection)
}
