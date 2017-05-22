
package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/standard"
)


func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Restricted Access
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"http://localhost:3000", "https://butterfli.io"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// ROUTES
	e.GET("/", accessible)
	r.GET("", restricted)


	e.POST("/login", LoginUserController)

	e.POST("/user", CreateUserController)
	e.GET("/:username", GetUserController)
	e.POST("/:username/update", UpdateUserController)
	e.POST("/:username/delete", RemoveUserController)

	
	e.POST("/:username/accounts/create", CreateAccountController)
	e.GET("/:username/accounts/:title", GetAccountController)
	e.POST("/:username/accounts/update/:title", UpdateAccountController)
	e.POST("/:username/accounts/delete/:title", DeleteAccountController)


	e.GET("/:username/accounts/:title/account-creds", FindAccountCredsController)
	e.POST("/:username/accounts/:title/account-creds", UpdateAccountCredsController)


	e.POST("/:username/accounts/:acctTitle/search/:socialNetwork", SearchController)
	e.POST("/:username/accounts/:acctTitle/favorite/:socialNetwork", SearchAndFavorite)


	e.GET("/:accountId/posts/:postId", FindPostController)
	// e.POST("/post/approve/:postId", ApprovePostController)

	// e.GET("/:username/accounts", GetAllAccountsByUsername)
	// e.GET("/users", GetAllUsers)
	// e.GET("/:username/accounts/:accountId/search/:socialNetwork/:searchTerm", SearchController)



	// e.POST("/user", CreateUser)


	// e.POST("/post/edit/:postId/title/:title", EditPost)
	// e.POST("/post/disapprove/:postId", DisapprovePost)
	// e.POST("/:username/accounts/:account_id/post/delete/:postId", RemovePost)
	// e.POST("/:username/accounts/:account_id/post/:postId/upload/twitter/:tweetText", PostTweet)
	// e.POST("/:username/accounts/:accountId/twitter/creds", UpdateAccountCredsController)

	// e.POST("/:username/botnet/favorite/:tweetId/accounts/:accountsArray", BotnetFavoriteTweet)
	// e.POST("/:username/botnet/follow-account/:accountId/accounts/:accountsArray", BotnetFollowAccountId)
	// e.POST("/:username/botnet/follow/:accountName/accounts/:accountsArray", BotnetFollowAccountName)

	// e.GET("/:username/accounts/:accountId/search-terms", GetAllSearchTerms)



	// // NOT TESTED
	// e.POST("/:username/accounts/create/:title", CreateAccount)


	fmt.Println("Server now running on port: 1323")
	e.Run(standard.New(":1323"))
}
