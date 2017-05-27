package store

import (
	"labix.org/v2/mgo"
	"os"
	"fmt"
)

func ConnectToDb() (*mgo.Session, error) {
	hostname := os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	port := os.Getenv("MONGO_PORT_27017_TCP_PORT")
	url := hostname + ":" + port
	session, err := mgo.Dial(url)
	if err != nil {panic(err)}

	return session, err
}

func ConnectToCollection(session *mgo.Session, collection_str string, keyString []string) (*mgo.Collection, error) {
	collection := session.DB("butterfli").C(collection_str)

	index := mgo.Index{
		Key:        keyString,
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := collection.EnsureIndex(index)
	if err != nil {
		fmt.Print("\nThere was an error in connecting to collection!\n")
		fmt.Print(err)
	}

	return collection, err
}
