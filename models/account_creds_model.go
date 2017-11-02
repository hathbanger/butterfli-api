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
	Id 					bson.ObjectId       `json:"id",bson:"_id,omitempty"`
	Timestamp 			time.Time	       	`json:"time",bson:"time,omitempty"`
	Username			string           	`json:"username",bson:"username,omitempty"`
	Account				string           	`json:"account",bson:"account,omitempty"`
	ConsumerKey			string           	`json:"consumerkey",bson:"consumerkey,omitempty"`
	ConsumerSecret		string           	`json:"consumersecret",bson:"consumersecret,omitempty"`
	AccessToken			string           	`json:"accesstoken",bson:"accesstoken,omitempty"`
	AccessTokenSecret	string          	`json:"accesstokensecret",bson:"accesstokensecret,omitempty"`
}


func NewAccountCreds(username string, account string) *AccountCreds {

	fmt.Println("New ACCOUNT CREDS")

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
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(
		session, "accountCreds", []string{"account"})
	if err != nil {
		panic(err)
	}

	accountCreds := &AccountCreds{
		Id: a.Id,
		Timestamp: a.Timestamp,
		Username: a.Username,
		Account: a.Account}

	err = collection.Insert(accountCreds)

	fmt.Print("accountCreds in new accountCreds", accountCreds)

	return nil
}



func FindAccountCredsModel(id string) (*AccountCreds, error) {

	session, err := store.ConnectToDb()
	defer session.Close()
	collection, err := store.ConnectToCollection(
		session, "accountCreds", []string{"account"})

	account, err := FindAccountModelId(id)

	accountCreds := &AccountCreds{}
	
	err = collection.Find(
		bson.M{"id": bson.ObjectIdHex(account.AccountCreds.Id.Hex())}).One(&accountCreds)
	if err != nil {
		fmt.Print(err)
	}

	return accountCreds, err
}


func UpdateAccountCredsModel(
	username string,
	title string,
	consumerKey string,
	consumerSecret string,
	accessToken string,
	accessTokenSecret string) (*AccountCreds, error) {

	account, err := FindAccountModelId(title)
	session, err := store.ConnectToDb()
	creds := account.AccountCreds
	defer session.Close()

	collection, err := store.ConnectToCollection(
		session, "accountCreds", []string{"account"})
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("\ncreds", creds, "\nconsumkey", consumerKey)

	colQuerier := bson.M{"id": creds.Id}
	change := bson.M{
		"$set": bson.M{ "consumerkey": newVar(
			creds.ConsumerKey, consumerKey),
			"consumersecret": newVar(
				creds.ConsumerSecret, consumerSecret),
			"accesstoken": newVar(
				creds.AccessToken, accessToken),
			"accesstokensecret": newVar(
				creds.AccessTokenSecret,
				accessTokenSecret)}}
	err = collection.Update(colQuerier, change)

	account, err = FindAccountModel(username, title)
	creds = account.AccountCreds
	fmt.Println("ACCT CREDS", creds)
	if err != nil {
		fmt.Print(err)
	}

	return creds, err
}

func newVar(a string, b string) string {
    if b != "" {
        return b
    }

    return a
}


func DeleteAccountCredsModel(id string) error {

	session, err := store.ConnectToDb()
	defer session.Close()
	collection, err := store.ConnectToCollection(
		session, "accountCreds", []string{"account"})
	
	err = collection.Remove(bson.M{"id": bson.ObjectIdHex(id)})
	if err != nil {
		fmt.Print(err)
	}
	
	return err
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