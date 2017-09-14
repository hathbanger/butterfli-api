package server


import (
	"fmt"
	"github.com/hathbanger/butterfli-api/models"
	// "github.com/labstack/echo"
	"github.com/ChimeraCoder/anaconda"
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
