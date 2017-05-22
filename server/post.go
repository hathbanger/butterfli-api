package server

import (
	// "github.com/labstack/echo"
	// "net/http"
	"github.com/hathbanger/butterfli-api/models"
	"github.com/labstack/echo"
	"net/http"
	"fmt"
	// "io/ioutil"
	// "encoding/base64"
	// "net/url"
	// "strconv"
	"strings"
	"github.com/ChimeraCoder/anaconda"
	// "log"
)


func CreatePostFromResults(username string, accountTitle string, searchTerm string, socialNetwork string, results anaconda.SearchResponse) anaconda.SearchResponse {
	var sinceTweetId = int64(0)
	var count = 0
	for _, tweet := range results.Statuses {
		if len(tweet.Entities.Media) > 0 {
			count = count + 1
			imgurl :=  tweet.Entities.Media[0].Media_url
			sinceTweetId = tweet.Id
			strArr := strings.Split(tweet.Text, " ")
			fmt.Printf("%q\n", strArr[:len(strArr) - 1])
			fmt.Print("%q\n", len(strArr))
			post := models.NewPost(username, accountTitle, tweet.Text, sinceTweetId, imgurl)
			err := post.Save()
			if err != nil {
				fmt.Print(" - - > Duplicate! \n")
			}
		}
	}
	// models.UpdateSearchTerm(searchTerm, sinceTweetId)
	// models.AddPostCountToSearchTerm(searchTerm, count)

	return results
}

func FindPostController(c echo.Context) error {
	accountId := c.Param("accountId")
	postId := c.Param("postId")
	post, err := models.FindPostById(accountId, postId)
	if err != nil {
		return c.JSON(http.StatusNotFound, "not found")
	}
	return c.JSON(http.StatusOK, post)
}

// func FindAllAccountPosts(c echo.Context) error {
// 	accountId := c.Param("account_id")
// 	posts, err := models.GetAllAccountPosts(accountId)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return c.JSON(http.StatusOK, posts)
// }

// func EditPost(c echo.Context) error {
// 	postId := c.Param("postId")
// 	title := c.Param("title")
// 	err := models.EditPostTitleById(postId, title)
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, "not approved")
// 	}
// 	return c.JSON(http.StatusOK, "approved!")
// }



// func DisapprovePost(c echo.Context) error {
// 	postId := c.Param("postId")
// 	err := models.DisapprovePostById(postId)
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, "failure")
// 	}
// 	return c.JSON(http.StatusOK, "disapproved!")
// }


// func RemovePost(c echo.Context) error {
// 	postId := c.Param("postId")
// 	err := models.DeletePost(postId)
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, "not able to remove the post..")
// 	}

// 	return c.JSON(http.StatusOK, "worked!!")
// }

// func PostTweet(c echo.Context) error {
// 	accountId := c.Param("account_id")
// 	postId := c.Param("postId")
// 	tweetText := c.Param("tweetText")
// 	fmt.Print("tweetText!")
// 	fmt.Print(tweetText)

// 	results := PostMediaToTwitter(accountId, postId, tweetText)

// 	return c.JSON(http.StatusOK, results)
// }

// func PostMediaToTwitter(accountId string, postId string, text string) anaconda.Tweet {
// 	post, err := models.FindPostById(postId)
// 	res, err := http.Get(post.Imgurl)
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	if err != nil {
// 		log.Fatalf("http.Get -> %v", err)
// 	}
// 	data, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatalf("ioutil.ReadAll -> %v", err)
// 	}
// 	api := AuthTwitter(accountId)
// 	mediaResponse, err := api.UploadMedia(base64.StdEncoding.EncodeToString(data))
// 	if err != nil {panic(err)}

// 	v := url.Values{}
// 	v.Set("media_ids", strconv.FormatInt(mediaResponse.MediaID, 10))
// 	res.Body.Close()

// 	fmt.Print("WOOO TEXT!")

// 	fmt.Print(text)

// 	u, err := url.QueryUnescape(text)

// 	result, err := api.PostTweet(u, v)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}
// 	return result
// }