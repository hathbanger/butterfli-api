package models

import (
	"time"
	// "fmt"

	"labix.org/v2/mgo/bson"
	"github.com/hathbanger/butterfli-api/store"
	// "github.com/butterfli-api/models"
	//"github.com/labstack/gommon/log"

)

type User struct {
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Username	string           `json:"username",bson:"username,omitempty"`
	Password	string           `json:"-",bson:"password,omitempty"`
	Accounts 	[]*Account 		`json:"accounts",bson:"accounts,omitempty"`
}


func NewUserModel(username string, password string) *User {
	u := new(User)
	u.Id = bson.NewObjectId()
	u.Username = username
	u.Password = password

	return u
}

func (u *User) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "users", []string{"username"})
	if err != nil {panic(err)}

	err = collection.Insert(&User{
		Id: u.Id,
		Timestamp: u.Timestamp,
		Username: u.Username,
		Password: u.Password})
	if err != nil {return err}

	return nil
}

func FindUserModel(username string) (User, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "users", []string{"username"})
	if err != nil {panic(err)}

	user := User{}
	err = collection.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return user, err
	}

	return user, err
}


func UpdateUserModel(username string, password string) (User, error) {
	user, err := FindUserModel(username)
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}
	collection := session.DB("butterfli").C("users")
	colQuerier := bson.M{"id": user.Id}
	change := bson.M{"$set": bson.M{ "password": password }}
	err = collection.Update(colQuerier, change)
	if err != nil {panic(err)}

	return user, err
}

func DeleteUserModel(username string) error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToCollection(session, "users", []string{"username"})
	if err != nil {panic(err)}

	err = collection.Remove(bson.M{"username": username})
	if err != nil {
		panic(err)
	}
	return nil
}
