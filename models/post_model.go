package models

import (
	"labix.org/v2/mgo/bson"
	"time"
	"github.com/butterfli-api/store"
)


type Post struct {
	//BaseModel
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Username	string           `json:"username",bson:"username,omitempty"`
	Account		string           `json:"account",bson:"account,omitempty"`
	Imgurl		string           `json:"imgurl",bson:"imgurl,omitempty"`
	Title		string           `json:"title",bson:"title,omitempty"`
	// SearchTerm	SearchTerm           `json:"searchterm",bson:"searchterm,omitempty"`
	OGSourceId	int64		`json:"ogSourceId",bson:"ogSourceId,omitempty"`
	Approved	bool           `json:"approved",bson:"approved,omitempty"`
	Rated		bool           `json:"rated",bson:"rated,omitempty"`
}

func NewPost(username string, accountTitle string, title string, ogSourceId int64, imgUrl string) *Post {
	p := new(Post)
	p.Id = bson.NewObjectId()
	p.Timestamp = time.Now()
	p.Username = username
	p.Account = accountTitle
	p.OGSourceId = ogSourceId
	p.Imgurl = imgUrl
	p.Title = title
	p.Approved = false
	p.Rated = false

	return p
}

func (p *Post) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToCollection(session, "posts", []string{"account", "imgurl"})
	if err != nil {
		panic(err)
	}
	post := &Post{
		Id: p.Id,
		Timestamp: p.Timestamp,
		Username: p.Username,
		Account: p.Account,
		Imgurl: p.Imgurl,
		Title: p.Title,
		Approved: p.Approved,
		OGSourceId: p.OGSourceId,
		Rated: p.Rated}

	err = collection.Insert(post)
	if err != nil {
		return err
	}


	collection, err = store.ConnectToCollection(session, "accounts", []string{"username", "title"})
	if err != nil {panic(err)}

	err = collection.Update(bson.M{"username": p.Username, "title": p.Title}, bson.M{"$push": bson.M{"posts": post}})

	if err != nil {
		return  err
	}

	return nil
}


// func FindPostById(post_id string) (*Post, error) {
// 	session, err := store.ConnectToDb()
// 	defer session.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	collection, err := store.ConnectToCollection(session, "posts", []string{"account", "imgurl"})
// 	if err != nil {
// 		panic(err)
// 	}
// 	post := Post{}
// 	err = collection.Find(bson.M{"id": bson.ObjectIdHex(post_id)}).One(&post)
// 	if err != nil {
// 		return &post, err
// 	}
// 	return &post, err
// }

// func GetAllAccountPosts(accountId string) ([]*Post, error){
// 	session, err := store.ConnectToDb()
// 	defer session.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	collection := session.DB("test").C("posts")
// 	posts := []*Post{}
// 	whereString := "this.rated == false || this.approved == true"
// 	err = collection.Find(bson.M{"$where": whereString, "account": accountId}).All(&posts)
// 	return posts, err
// }

// func DeletePost(id string) error {
// 	session, err := store.ConnectToDb()
// 	collection := session.DB("test").C("posts")
// 	err = collection.Remove(bson.M{"id": bson.ObjectIdHex(id)})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return nil
// }


// func EditPostTitleById(postId string, title string) error {
// 	session, err := store.ConnectToDb()
// 	collection := session.DB("test").C("posts")
// 	colQuerier := bson.M{"id": bson.ObjectIdHex(postId)}
// 	change := bson.M{"$set": bson.M{ "title": title }}
// 	err = collection.Update(colQuerier, change)
// 	if err != nil {panic(err)}

// 	return nil
// }

// func ApprovePostById(postId string) error {
// 	session, err := store.ConnectToDb()
// 	collection := session.DB("test").C("posts")
// 	colQuerier := bson.M{"id": bson.ObjectIdHex(postId)}
// 	change := bson.M{"$set": bson.M{ "approved": true, "rated": true }}
// 	err = collection.Update(colQuerier, change)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return nil
// }

// func DisapprovePostById(postId string) error {
// 	session, err := store.ConnectToDb()
// 	collection := session.DB("test").C("posts")
// 	colQuerier := bson.M{"id": bson.ObjectIdHex(postId)}
// 	change := bson.M{"$set": bson.M{ "approved": false, "rated": true }}
// 	err = collection.Update(colQuerier, change)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return nil
// }

