package server


import (
	"github.com/hathbanger/butterfli-api/models"
	"github.com/labstack/echo"
	"github.com/ChimeraCoder/anaconda"
	"net/http"
	"fmt"

	"net/url"
)


func SearchController(c echo.Context) error {

	socialNetwork := c.Param("socialNetwork")
	// searchTermString := "catsareassholes"
	acctTitle := c.Param("acctTitle")
	username := c.Param("username")
	searchTermString := c.FormValue("searchTerm")
	// searchTerm, err := models.FindSearchTerm(acctTitle, searchTermString)
	// if err != nil {
	// 	searchTerm = models.NewSearchTerm(accountId, searchTermString)
	// 	searchTerm.Save()
	// }
	fmt.Println("searchTermString", searchTermString)
	results := Search(username, acctTitle, socialNetwork, searchTermString)
	CreatePostFromResults(
		username, acctTitle, searchTermString, socialNetwork, results)

	return c.JSON(http.StatusOK, results)
}


func Search(
	username string,
	acctTitle string,
	socialNetwork string,
	searchTerm string) anaconda.SearchResponse {

	switch socialNetwork {
	case "twitter-img":

		return SearchTwitter(
			username,
			acctTitle,
			searchTerm,
			"10",
			" filter:twimg",
		)
	default:
		panic("unrecognized escape character")
	}
}


func SearchTwitter(
	username string,
	acctTitle string,
	searchTerm string,
	count string,
	searchType string) anaconda.SearchResponse {

	v := url.Values{}
	// s := strconv.FormatInt(searchTerm.SinceTweetId, 10)
	// v.Set("since_id", s)
	v.Add("count", count)
	updatedSearch := searchTerm + searchType
	api := AuthTwitter(username, acctTitle)
	search_result, err := api.GetSearch(updatedSearch, v)
	if err != nil {
		panic(err)
	}

	return search_result
}

func SearchAndFavorite(c echo.Context) error {

	fmt.Println("WOWWWW")
	acctTitle := c.Param("acctTitle")
	username := c.Param("username")
	searchTermString := c.FormValue("searchTerm")

	fmt.Println(acctTitle)

	// searchTermString := c.Param("searchTerm")
	accountId := c.Param("accountId")

	api := AuthTwitter(username, acctTitle)
	searchTerm, err := models.FindSearchTerm(accountId, searchTermString)

	fmt.Print("WOAH! NEW TERM\n\n")
	fmt.Print(err)
	fmt.Print("That was an Err\n\n")
	fmt.Print(searchTerm)
	fmt.Print("\n this was a searchterm\n\n")

	// if searchTerm != nil {
	// 	searchTerm = models.NewSearchTerm(accountId, searchTermString)
	// 	searchTerm.Save()

	// 	fmt.Print("\n\nSAVEDDDD! NEW TERM\n\n")
	// 	fmt.Print(err)
	// 	fmt.Print(searchTerm.Text)
	// 	fmt.Print("\nNEW TERM\n\n")
	// }

	//results := Search(username, accountId, socialNetwork, *favoriteTerm)
	v := url.Values{}
	// s := strconv.FormatInt(favoriteTerm.SinceTweetId, 10)
	// v.Set("since_id", s)
	v.Add("count", "10")
	// updatedSearch := favoriteTerm.Text
	search_result, err := api.GetSearch(searchTermString, v)
	if err != nil {
		panic(err)
	}

	var successes = 0
	var failures = 0
	for _, tweet := range search_result.Statuses {
		res, err := api.Favorite(tweet.Id)
		// models.UpdateFavoriteTerm(favoriteTerm, tweet.Id)
		if res.Id != 0  {
			successes = successes + 1
			fmt.Print(" Success!")
		}

		if err != nil {
			failures = failures + 1
			fmt.Print("error!")
			fmt.Print(err)
		}
	}
	fmt.Print(successes)

	return c.JSON(
		http.StatusOK,
		fmt.Sprintf(
			"AccountId %s just favorited %v new tweets, and failed %v times",
			accountId,
			successes,
			failures))
}


//func FavoriteTwitter(username string, accountId string, searchTerm models.SearchTerm) anaconda.SearchResponse {
//	v := url.Values{}
//	s := strconv.FormatInt(searchTerm.SinceTweetId, 10)
//	v.Set("since_id", s)
//	v.Add("count", "30")
//	updatedSearch := searchTerm.Text + " filter:twimg"
//	api := AuthTwitter(accountId)
//	search_result, err := api.GetSearch(updatedSearch, v)
//	if err != nil {panic(err)}
//
//	fmt.Print("search_result:")
//	fmt.Print(search_result)
//
//	return search_result
//}


// func FavoriteTweets(c echo.Context) error {

// }



// func GetAllSearchTerms(c echo.Context) error {
// 	accountId := c.Param("accountId")
// 	searchTerms := models.FindAllSearchTerms(accountId)
// 	return c.JSON(http.StatusOK, searchTerms)
// }

// func BotnetFavoriteTweet(c echo.Context) error {
// 	tweetId := c.Param("tweetId")
// 	accountsArray := c.Param("accountsArray")
// 	tweetId64, err := strconv.ParseInt(tweetId, 10, 64)
// 	if err != nil {
// 		panic(err)
// 	}
// 	accountsSlice := strings.Split(accountsArray, "+")
// 	for _, accountId := range accountsSlice {
// 		api := AuthTwitter(accountId)
// 		api.Favorite(tweetId64)
// 		api.EnableThrottling(10*time.Second, 5)
// 	}

// 	return c.JSON(http.StatusOK, fmt.Sprintf("Sick! You just liked tweetId %s with %v accounts", tweetId, len(accountsSlice)))
// }


// func BotnetFollowAccountId(c echo.Context) error {
// 	followAccountId := c.Param("accountId")
// 	accountsArray := c.Param("accountsArray")
// 	followAccountId64, err := strconv.ParseInt(followAccountId, 10, 64)
// 	if err != nil {
// 		panic(err)
// 	}
// 	accountsSlice := strings.Split(accountsArray, "+")
// 	for _, accountId := range accountsSlice {
// 		api := AuthTwitter(accountId)
// 		api.FollowUserId(followAccountId64, nil)
// 		api.EnableThrottling(10*time.Second, 5)
// 	}

// 	return c.JSON(http.StatusOK, fmt.Sprintf("Sick! You just followed accountId %s with %v accounts", followAccountId, len(accountsSlice)))
// }

// func BotnetFollowAccountName(c echo.Context) error {
// 	followAccountName := c.Param("accountName")
// 	accountsArray := c.Param("accountsArray")

// 	accountsSlice := strings.Split(accountsArray, "+")
// 	for _, accountId := range accountsSlice {
// 		api := AuthTwitter(accountId)
// 		api.FollowUser(followAccountName)
// 		api.EnableThrottling(10*time.Second, 5)
// 	}

// 	return c.JSON(http.StatusOK, fmt.Sprintf("Sick! You just followed %s with %v accounts", followAccountName, len(accountsSlice)))
// }