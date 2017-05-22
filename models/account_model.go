package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	"github.com/hathbanger/butterfli-api/store"
	// "labix.org/v2/mgo"
	"fmt"
)

type Account struct {
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Title		string           	`json:"title",bson:"title,omitempty"`
	Username	string           	`json:"username",bson:"username,omitempty"`
	AccountCreds    *AccountCreds		`json:"accountcreds",bson:"accountcreds,omitempty"`
	Posts 			[]*Post		 	`json:"posts",bson:"posts,omitempty"`
	// SearchTerms 	[]*SearchTerm		 `json:"searchterms",bson:"searchterms,omitempty"`
	// FavoriteTerms 	[]*FavoriteTerm		 `json:"favoriteterms",bson:"favoriteterms,omitempty"`
}

func NewAccountModel(username string, title string) *Account {
	a := new(Account)
	a.Id = bson.NewObjectId()
	a.Timestamp = time.Now()
	a.Username = username
	a.Title = title
	a.AccountCreds = NewAccountCreds(username, title)

	return a
}

func (a *Account) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToCollection(session, "accounts", []string{"username", "title"})
	if err != nil {
		panic(err)
	}
	a.AccountCreds.Save()
	err = collection.Insert(&Account{
		Id: a.Id,
		Timestamp: a.Timestamp,
		Title: a.Title,
		Username: a.Username,
		AccountCreds: a.AccountCreds,
	})



	collection, err = store.ConnectToCollection(session, "users", []string{"username"})
	if err != nil {panic(err)}

	err = collection.Update(bson.M{"username": a.Username}, bson.M{"$push": bson.M{"accounts": a}})

	if err != nil {
		return  err
	}


	return nil
}


func FindAccountModel(username string, title string) (*Account, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToCollection(session, "accounts", []string{"title", "username"})
	if err != nil {
		panic(err)
	}
	account := Account{}
	err = collection.Find(bson.M{"username": username, "title": title}).One(&account)
	if err != nil {
		panic(err)
	}
	return &account, err
}


func FindAccountModelId(accountId string) (*Account, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToCollection(session, "accounts", []string{"title", "username"})
	if err != nil {
		panic(err)
	}
	account := Account{}
	err = collection.Find(bson.M{"id": bson.ObjectIdHex(accountId)}).One(&account)
	if err != nil {
		panic(err)
	}
	return &account, err
}


func UpdateAccountModel(username string, oldTitle string, newTitle string) (*Account, error) {
	account, err := FindAccountModel(username, oldTitle)
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}
	collection := session.DB("butterfli").C("accounts")
	colQuerier := bson.M{"id": account.Id}
	change := bson.M{"$set": bson.M{ "title": newTitle }}
	err = collection.Update(colQuerier, change)
	if err != nil {panic(err)}

	return account, err
}


func DeleteAccountModel(username string, title string) error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		fmt.Print(err)
	}	
	collection, err := store.ConnectToCollection(session, "accounts", []string{"title", "username"})
	if err != nil {
		fmt.Print(err)
	}
	err = collection.Remove(bson.M{"title": title, "username": username})
	if err != nil {
		fmt.Print(err)
	}
	return err
}



// func FindAccountById(account_id string) (*Account, error) {
// 	session, err := store.ConnectToDb()
// 	defer session.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	collection := ConnectAccounts(session)
// 	if err != nil {
// 		panic(err)
// 	}
// 	account := Account{}
// 	err = collection.Find(bson.M{"id": bson.ObjectIdHex(account_id)}).One(&account)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return &account, err
// }

// func (u *User) FindAccountByTitle(title string) (*Account, error) {
// 	session, err := store.ConnectToDb()
// 	defer session.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	collection, err := store.ConnectToCollection(session, "accounts", []string{"username", "title"})
// 	if err != nil {
// 		//panic(err)
// 		return &Account{}, err
// 	}
// 	account := Account{}
// 	err = collection.Find(bson.M{"username": u.Username, "title": title}).One(&account)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return &account, err
// }

// func GetAllAccounts(username string) ([]*Account, error){
// 	session, err := store.ConnectToDb()
// 	defer session.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	collection := ConnectAccounts(session)
// 	accounts := []*Account{}
// 	err = collection.Find(bson.M{"username": username}).All(&accounts)
// 	return accounts, err
// }


// func ConnectAccounts(session *mgo.Session) *mgo.Collection{
// 	collection, err := store.ConnectToCollection(session, "accounts", []string{"username", "title"})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return collection
// }