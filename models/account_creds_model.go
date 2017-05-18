package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/butterfli-api/store"
	"time"
	// "fmt"
	//"go/doc"
	//"labix.org/v2/mgo"
	//"go/doc"
)

type AccountCreds struct {
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Username	string           `json:"username",bson:"username,omitempty"`
	Account		string           `json:"account",bson:"account,omitempty"`
	ConsumerKey		string           `json:"-",bson:"consumerKey,omitempty"`
	ConsumerSecret		string           `json:"-",bson:"consumerSecret,omitempty"`
	AccessToken		string           `json:"-",bson:"accessToken,omitempty"`
	AccessTokenSecret		string           `json:"-",bson:"accessTokenSecret,omitempty"`
}

func NewAccountCreds(username string, account string) *AccountCreds {
	a := new(AccountCreds)
	a.Id = bson.NewObjectId()
	a.Timestamp = time.Now()
	a.Username = username
	a.Account = account

	return a
}

func (a *AccountCreds) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "accountCreds", []string{"account", "username"})
	if err != nil {panic(err)}

	accountCreds := &AccountCreds{
		Id: a.Id,
		Timestamp: a.Timestamp,
		Username: a.Username,
		Account: a.Account}

	err = collection.Insert(accountCreds)

	return nil
}

// func UpdateAccountCreds(accountCreds string, consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) error {
// 	session, err := store.ConnectToDb()
// 	//if err != nil {panic(err)}


// 	collection, err := store.ConnectToCollection(session, "accountCreds", []string{"account"})
// 	//if err != nil {panic(err)}




// 	//acct, err := FindAccountCredsById(accountCreds)

// 	fmt.Print("accountCreds! ")
// 	fmt.Print(bson.IsObjectIdHex(accountCreds))



// 	colQuerier := bson.M{"_id": bson.ObjectIdHex(accountCreds)}
// 	change := bson.M{"$set": bson.M{ "consumerKey": consumerKey }}
// 	if err = collection.Update(colQuerier, change); err != nil {
// 		fmt.Print(err)
// 	}


// 	if err != nil {
// 		fmt.Print("\nissssues!\n")
// 	}




// 	return err
// }


// func FindAccountCredsById(accountCredsId string) (*AccountCreds, error) {
// 	session, err := store.ConnectToDb()
// 	defer session.Close()
// 	if err != nil {panic(err)}

// 	collection, err := store.ConnectToCollection(session, "accountCreds", []string{"imgurl"})
// 	if err != nil {panic(err)}

// 	accountCreds := AccountCreds{}
// 	err = collection.Find(bson.M{"id": bson.ObjectIdHex(accountCredsId)}).One(&accountCreds)
// 	if err != nil {panic(err)}

// 	return &accountCreds, err
// }


// func FindAccountCredsByAccountId(accountId string) (*AccountCreds, error) {
// 	session, err := store.ConnectToDb()
// 	defer session.Close()
// 	if err != nil {panic(err)}

// 	fmt.Print("bout to do FindAccountCredsByAccountId!")
// 	collection, err := store.ConnectToCollection(session, "accountCreds", []string{"account"})
// 	if err != nil {
// 		fmt.Print("error finding collection")
// 		fmt.Print(err)
// 	}

// 	accountCreds := AccountCreds{}
// 	err = collection.Find(bson.M{"account": accountId}).One(&accountCreds)
// 	if err != nil {
// 		fmt.Print("error finding account creds")
// 		fmt.Print(err)
// 	}

// 	return &accountCreds, err
// }