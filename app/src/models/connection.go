package models

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

//DatabaseHost ...
var DatabaseHost = "mongo:27017"

//Database session
type Database struct {
	db *mgo.Database
}

//Connect Returns a database session
func Connect() Database {
	session, err := mgo.Dial(DatabaseHost)

	if err != nil {
		fmt.Println("Failed to stablish connection:", err)
		return Database{nil}
	}

	return Database{session.DB("yawoen")}
}

//Close current session
func (d Database) Close() {
	if d.db != nil {
		d.db.Session.Close()
	}
}

//Collection Access collection object of a database
func (d Database) Collection(collection string) *mgo.Collection {
	if d.db == nil {
		return nil
	}

	return d.db.C(collection)
}
