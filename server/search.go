package server


import (
	"github.com/hathbanger/butterfli-api/models"
	"github.com/labstack/echo"
	"github.com/ChimeraCoder/anaconda"

	"github.com/dghubble/go-twitter/twitter"




	"net/http"
	"time"
	"fmt"
	"strconv"
	// "sync"
	"net/url"
)


func SearchController(c echo.Context) error {
	socialNetwork := c.Param("socialNetwork")
	acctTitle := c.Param("acctTitle")
	username := c.Param("username")
	searchTermString := c.FormValue("searchTerm")
	searchTerm, err := models.FindSearchTerm(acctTitle, searchTermString)
	results := Search(username, acctTitle, socialNetwork, searchTerm)
	CreatePostFromResults(
		username, acctTitle, searchTerm, socialNetwork, results)

	fmt.Println(searchTerm, err)
	return c.JSON(http.StatusOK, results)
}


func Search(
	username string,
	acctTitle string,
	socialNetwork string,
	searchTerm *models.SearchTerm) anaconda.SearchResponse {

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
	searchTerm *models.SearchTerm,
	count string,
	searchType string) anaconda.SearchResponse {

	v := url.Values{}
	s := strconv.FormatInt(searchTerm.SinceTweetId, 10)
	v.Set("since_id_str", s)
	v.Set("count", count)
	updatedSearch := searchTerm.Text + searchType
	api := AuthTwitter(username, acctTitle)
	search_result, err := api.GetSearch(updatedSearch, v)
	fmt.Println("\n\tWOOO: ", v)
	if err != nil {
		panic(err)
	}

	return search_result
}

func SearchAndFavorite(c echo.Context) error {
	username := c.Param("username")
	searchTermString := c.FormValue("searchTerm")
	accountId := c.Param("acctTitle")
	count := c.FormValue("count")

	client := AuthTwitterClient(username, accountId)
	i, _ := strconv.Atoi(count)


	searchTerm, _ := models.FindSearchTerm(accountId, searchTermString)

	searchParams := &twitter.SearchTweetParams{
		Query:      searchTermString,
		Count:      i,
		ResultType: "recent",
		Lang:       "en",
		SinceID: searchTerm.SinceTweetId,
	}	

	var successes = 0
	var failures = 0
	var tweetId int64


	searchResult, _, _ := client.Search.Tweets(searchParams)
	for _, tweet := range searchResult.Statuses {
		favoriteParams := &twitter.FavoriteCreateParams{
			ID:      tweet.ID,
		}
		_, http, _ := client.Favorites.Create(favoriteParams)
		if http.StatusCode == 200 {
			successes = successes + 1
			tweetId = favoriteParams.ID
		}

		if http.StatusCode != 200 {
			failures = failures + 1
		}

		time.Sleep(time.Second * 12)
		fmt.Printf("Favorited: %+v\n\n", tweet.Text)
	}

	models.UpdateSearchTerm(searchTerm, tweetId)
	fmt.Print(searchTerm)


	fmt.Print("successes:", successes)
	fmt.Print("failures:", failures, "\n")
	fmt.Print("searchTerm:", searchTerm, "\n")

	return c.JSON(
		http.StatusOK,
		fmt.Sprintf(
			"\n\nAccountId %s just favorited %v new tweets, and failed %v times\n\n",
			accountId,
			successes,
			failures))
}

func UnfavoriteTweetsLoop(c echo.Context, count int, successes int, failures int) (int, int) {
	username := c.Param("username")
	accountId := c.Param("acctTitle")

	client := AuthTwitterClient(username, accountId)
	searchParams := &twitter.FavoriteListParams{
		Count: count,
	}


	fmt.Print("\n\nLoop starting\n")
	
	
	listResults, _, _ := client.Favorites.List(searchParams)

	resLen := len(listResults)
	fmt.Print(count)
	fmt.Print(resLen)

	if resLen == 0 {return successes, failures}

	for _, tweet := range listResults {
		unfavoriteParams := &twitter.FavoriteDestroyParams{
			ID:      tweet.ID,
		}
		_, http, _ := client.Favorites.Destroy(unfavoriteParams)
		
		if http.StatusCode == 200 {
			successes = successes + 1
		}
		if http.StatusCode != 200 {
			failures = failures + 1
		}

		time.Sleep(time.Second * 12)
		fmt.Printf("Destroyed: %+v\n\n", tweet.Text)
		fmt.Printf("successes: %+v", successes)
		fmt.Printf(" failures: %+v\n\n", failures)

		if count == successes { break }
	}

	if count > successes {
		return UnfavoriteTweetsLoop(c, count, successes, failures)
	}

	return successes, failures

}

func UnfavoriteTweets(c echo.Context) error {
	count := c.FormValue("count")
	var successes = 0
	var failures = 0
	stringCount, _ := strconv.Atoi(count)

	loopSuccess, loopFails := UnfavoriteTweetsLoop(c, stringCount, successes, failures)

	successes += loopSuccess
	failures += loopFails

	fmt.Print("successes:", successes)
	fmt.Print("failures:", failures, "\n")

	return c.JSON(
		http.StatusOK,
		fmt.Sprintf(
			"\n\nAccountId %s just unfavorited %v new tweets, and failed %v times\n\n",
			"accountId",
			"successfulUnfavorites",
			"failures"))	
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