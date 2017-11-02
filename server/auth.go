package server


import (
	"fmt"
	"github.com/hathbanger/butterfli-api/models"
	// "github.com/labstack/echo"
	"github.com/ChimeraCoder/anaconda"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"	
)

func AuthTwitter(username string, acctTitle string) *anaconda.TwitterApi {
	fmt.Print(acctTitle)
	accountCreds, err := models.FindAccountCredsModel(acctTitle)
	anaconda.SetConsumerKey(accountCreds.ConsumerKey)
	anaconda.SetConsumerSecret(accountCreds.ConsumerSecret)
	api := anaconda.NewTwitterApi(accountCreds.AccessToken, accountCreds.AccessTokenSecret)
	if err != nil {
		panic(err)
	}

	return api
}

func AuthTwitterClient(username string, acctTitle string) *twitter.Client {
	accountCreds, err := models.FindAccountCredsModel(acctTitle)
	config := oauth1.NewConfig(accountCreds.ConsumerKey, accountCreds.ConsumerSecret)
	token := oauth1.NewToken(accountCreds.AccessToken, accountCreds.AccessTokenSecret)

	fmt.Print(accountCreds.ConsumerKey)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	if err != nil {
		panic(err)
	}

	return client
}
