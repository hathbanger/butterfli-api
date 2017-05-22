package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/hathbanger/butterfli-api/store"
	"time"
	"fmt"
	//"go/doc"
	//"labix.org/v2/mgo"
	//"go/doc"
)

type AccountCreds struct {
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Username	string           `json:"username",bson:"username,omitempty"`
	Account		string           `json:"account",bson:"account,omitempty"`
	ConsumerKey		string           `json:"consumerkey",bson:"consumerkey,omitempty"`
	ConsumerSecret		string           `json:"consumersecret",bson:"consumersecret,omitempty"`
	AccessToken		string           `json:"accesstoken",bson:"accesstoken,omitempty"`
	AccessTokenSecret		string           `json:"accesstokensecret",bson:"accesstokensecret,omitempty"`
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

	collection, err := store.ConnectToCollection(session, "accounts", []string{"account"})
	if err != nil {panic(err)}

	accountCreds := &AccountCreds{
		Id: a.Id,
		Timestamp: a.Timestamp,
		Username: a.Username,
		Account: a.Account}

	err = collection.Insert(accountCreds)

	return nil
}



func FindAccountCredsModel(username string, title string) (*AccountCreds, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	collection, err := store.ConnectToCollection(session, "accounts", []string{"account"})
	
	account := Account{}
	err = collection.Find(bson.M{"username": username, "title": title}).One(&account)
	if err != nil {
		fmt.Print(err)
	}

	return account.AccountCreds, err
}


func UpdateAccountCredsModel(username string, title string, consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) (*AccountCreds, error) {
	account, err := FindAccountModel(username, title)
	session, err := store.ConnectToDb()
	creds := account.AccountCreds
	defer session.Close()

	collection, err := store.ConnectToCollection(session, "accounts", []string{"account"})


	if err != nil {fmt.Print(err)}

	colQuerier := bson.M{"id": account.Id}
	change := bson.M{"$set": bson.M{ "accountcreds.consumerkey": newVar(creds.ConsumerKey, consumerKey), "accountcreds.consumersecret": newVar(creds.ConsumerSecret, consumerSecret), "accountcreds.accesstoken": newVar(creds.AccessToken, accessToken), "accountcreds.accesstokensecret": newVar(creds.AccessTokenSecret, accessTokenSecret)}}
	err = collection.Update(colQuerier, change)

	account, err = FindAccountModel(username, title)
	creds = account.AccountCreds
	if err != nil {fmt.Print(err)}

	return creds, err
}




func newVar(a string, b string) string {
    if b != "" {
        return b
    }
    return a
}




// func UpdateAccountCreds(accountCreds string, consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) error {
// 	session, err := store.ConnectToDb()
// 	//if err != nil {panic(err)}


// 	collection, err := store.ConnectToCollection(session, "accounts", []string{"account"})
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


// func FindAccountCredsByAccountId(accountId string) (*AccountCreds, error) {
// 	session, err := store.ConnectToDb()
// 	defer session.Close()
// 	if err != nil {panic(err)}

// 	fmt.Print("bout to do FindAccountCredsByAccountId!")
// 	collection, err := store.ConnectToCollection(session, "accounts", []string{"account"})
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