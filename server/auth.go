package server


import (
	"github.com/hathbanger/butterfli-api/models"
	// "github.com/labstack/echo"
	"github.com/ChimeraCoder/anaconda"
)

func AuthTwitter(username string, acctTitle string) *anaconda.TwitterApi {
	accountCreds, err := models.FindAccountCredsModel(username, acctTitle)
	anaconda.SetConsumerKey(accountCreds.ConsumerKey)
	anaconda.SetConsumerSecret(accountCreds.ConsumerSecret)
	api := anaconda.NewTwitterApi(accountCreds.AccessToken, accountCreds.AccessTokenSecret)

	if err != nil {panic(err)}

	return api
}
