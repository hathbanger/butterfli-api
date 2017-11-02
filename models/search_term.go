package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/hathbanger/butterfli-api/store"
	"time"
	"fmt"
)

type SearchTerm struct {
	Id 		bson.ObjectId 		`json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       	`json:"time",bson:"time,omitempty"`
	Text		string           	`json:"text",bson:"text,omitempty"`
	Account		string           	`json:"account",bson:"account,omitempty"`
	PostCount	int           		`json:"postcount",bson:"postcount,omitempty"`
	SinceTweetId	int64 			`json:"sincetweetid",bson:"sincetweetid,omitempty"`
}

func NewSearchTerm(account string, text string) *SearchTerm {
	var sinceTweetId int64
	s := new(SearchTerm)
	s.Id = bson.NewObjectId()
	s.Text = text
	s.Account = account
	s.SinceTweetId = sinceTweetId

	return s
}

func (s *SearchTerm) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "searchTerms", []string{"account", "text"})
	if err != nil {panic(err)}

	searchTerm := &SearchTerm{
		Id: s.Id,
		Timestamp: s.Timestamp,
		Text: s.Text,
		Account: s.Account,
		SinceTweetId: s.SinceTweetId}

	err = collection.Insert(searchTerm)

	if err != nil {panic(err)}
	return nil

}

func FindAllSearchTerms(accountId string) []*SearchTerm {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "searchTerms", []string{"account", "text"})
	if err != nil {panic(err)}

	searchTerms := []*SearchTerm{}
	err = collection.Find(bson.M{"account": accountId}).All(&searchTerms)
	if err != nil {panic(err)}

	return searchTerms
}

func FindSearchTerm(account string, text string) (*SearchTerm, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "searchTerms", []string{"account", "text"})
	if err != nil {panic(err)}

	searchTerm := SearchTerm{}

	count, err := collection.Find(bson.M{ "account": account, "text": text}).Count()
	err = collection.Find(bson.M{ "account": account, "text": text}).One(&searchTerm)

	if count == 0 {
		newSearchTerm := NewSearchTerm(account, text)
		newSearchTerm.Save()
	}

	return &searchTerm, err
}


func UpdateSearchTerm(searchTerm *SearchTerm, sinceTweetId int64) error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}


	collection, err := store.ConnectToCollection(session, "searchTerms", []string{"account", "text"})
	if err != nil {panic(err)}

	fmt.Print(sinceTweetId)

	colQuerier := bson.M{"id": searchTerm.Id}
	change := bson.M{"$set": bson.M{"sincetweetid": sinceTweetId}}
	err = collection.Update(colQuerier, change)
	if err != nil {
		fmt.Print("\nissssues!\n")
	}

	fmt.Print("\nBOOM, added sincetweetid\n", sinceTweetId, "   booooommm\n\n")

	return err
}

func AddPostCountToSearchTerm(searchTerm *SearchTerm, count int) error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {panic(err)}

	collection, err := store.ConnectToCollection(session, "searchTerms", []string{"account", "text"})
	if err != nil {panic(err)}

	colQuerier := bson.M{"id": searchTerm.Id}
	change := bson.M{"$inc": bson.M{"postcount": count}}
	err = collection.Update(colQuerier, change)
	if err != nil {fmt.Print("\nissssues!\n")}

	return err
}