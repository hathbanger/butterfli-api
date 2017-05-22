package server

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/hathbanger/butterfli-api/models"
	//"fmt"
)

func CreateAccountController(c echo.Context) error {
	username := c.Param("username")
	title := c.FormValue("title")
	account := models.NewAccountModel(username, title)
	err := account.Save()
	if err != nil {
		return c.JSON(http.StatusForbidden, "We're sorry! We couldn't create an account for you.")
	}
	return c.JSON(http.StatusOK, account)
}


func GetAccountController(c echo.Context) error {
	username := c.Param("username")
	title := c.Param("title")
	account, err := models.FindAccountModel(username, title)
	if err != nil {
		return c.JSON(http.StatusForbidden, "We're sorry! We couldn't get the account for you.")
	}
	return c.JSON(http.StatusOK, account)
}


func UpdateAccountController(c echo.Context) error {
	username := c.Param("username")
	oldTitle := c.Param("title")
	newTitle := c.FormValue("title")
	models.UpdateAccountModel(username, oldTitle, newTitle)
	account, err := models.FindAccountModel(username, newTitle)
	if err != nil {
		return err
	}	
	return c.JSON(http.StatusOK, account)
}

func DeleteAccountController(c echo.Context) error {
	username := c.Param("username")
	title := c.Param("title")
	err := models.DeleteAccountModel(username, title)
	if err != nil {
		return c.JSON(http.StatusNotFound, "not able to remove the account..")
	}

	return c.JSON(http.StatusOK, "Account deleted!")	
}




func FindAccountCredsController(c echo.Context) error {
	username := c.Param("username")
	acctTitle := c.Param("title")
	accountCreds, err := models.FindAccountCredsModel(username, acctTitle)

	if err != nil {
		return c.JSON(http.StatusForbidden, "We're sorry! There was an issue finding your acct creds..")
	}

	return c.JSON(http.StatusOK, accountCreds)
}


func UpdateAccountCredsController(c echo.Context) error {
	username := c.Param("username")
	acctTitle := c.Param("title")
	newConsumerKey := c.FormValue("consumerKey")
	newConsumerSecret := c.FormValue("consumerSecret")
	newAccessToken := c.FormValue("accessToken")
	newAccessTokenSecret := c.FormValue("accessTokenSecret")
	accountCreds, err := models.UpdateAccountCredsModel(username, acctTitle, newConsumerKey, newConsumerSecret, newAccessToken, newAccessTokenSecret)

	if err != nil {
		return c.JSON(http.StatusForbidden, "We're sorry! There was an issue updating your account credentials..")
	}

	return c.JSON(http.StatusOK, accountCreds)
}



// func GetAllAccountsByUsername(c echo.Context) error {
// 	username := c.Param("username")
// 	accounts, err := models.GetAllAccounts(username)
// 	if err != nil {panic(err)}

// 	return c.JSON(http.StatusOK, accounts)
// }



// func RemoveAccount(c echo.Context) error {
// 	accountId := c.Param("account_id")
// 	err := models.DeleteAccount(accountId)
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, "not able to remove the account..")
// 	}

// 	return c.JSON(http.StatusOK, "worked!!")
// }


// func UpdateAccountCredsController(c echo.Context) error {
// 	//username := c.Param("username")
// 	accountId := c.Param("accountId")
// 	//accountCreds, err := models.FindAccountCredsByAccountId(accountId)
// 	consumerKey := c.FormValue("consumerKey")
// 	consumerSecret := c.FormValue("consumerSecret")
// 	accessToken := c.FormValue("accessToken")
// 	accessTokenSecret := c.FormValue("accessTokenSecret")
// 	err := models.UpdateAccountCreds(accountId, consumerKey, consumerSecret, accessToken, accessTokenSecret)
// 	//err := account
// 	//fmt.Print("\n\naccount creds update: \n\n\n")
// 	//fmt.Print(accountCreds)
// 	//fmt.Print(consumerKey)
// 	//fmt.Print(err)
// 	//fmt.Print(consumerSecret)
// 	//fmt.Print(accessToken)
// 	//fmt.Print(accessTokenSecret)
// 	if err != nil {
// 		return c.JSON(http.StatusForbidden, "We're sorry! There was an issue..")
// 	}

// 	return c.JSON(http.StatusOK, accountId)
// }


// func GetAccountCreds(c echo.Context) error {
// 	accountId := c.Param("accountId")
// 	account, err := models.FindAccountCredsByAccountId(accountId)
// 	if err != nil {
// 		return c.JSON(http.StatusForbidden, "We're sorry! we couldn't find it....")
// 	}
// 	return c.JSON(http.StatusOK, account)
// }




// func GetAccountById(c echo.Context) error {
// 	account_id := c.Param("account_id")
// 	account, err := models.FindAccountById(account_id)
// 	if err != nil {panic(err)}
// 	if account.Id != "" {
// 		return c.JSON(http.StatusOK, account)
// 	} else {
// 		return c.JSON(http.StatusNotFound, "not found")
// 	}
// }

