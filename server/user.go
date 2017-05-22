package server

import (
	"net/http"

	"time"
	"github.com/hathbanger/butterfli-api/models"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
)

func CreateUserController(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	user := models.NewUserModel(username, password)
	err := user.Save()
	if err != nil {
		return c.JSON(http.StatusForbidden, "We're sorry! There's already a user with that username..")
	}
	return c.JSON(http.StatusOK, user)
}


func LoginUserController(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	user, err := models.FindUserModel(username)
	if err != nil {panic(err)}
	if user.Password == password {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = user.Username
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func GetUserController(c echo.Context) error {
	username := c.Param("username")
	user, err := models.FindUserModel(username)
	if err != nil {
		return err
	}	
	return c.JSON(http.StatusOK, user)
}

func UpdateUserController(c echo.Context) error {
	username := c.Param("username")
	password := c.FormValue("password")
	models.UpdateUserModel(username, password)
	user, err := models.FindUserModel(username)
	if err != nil {
		return err
	}	
	return c.JSON(http.StatusOK, user)
}


func RemoveUserController(c echo.Context) error {
	username := c.Param("username")
	err := models.DeleteUserModel(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, "not able to remove the account..")
	}

	return c.JSON(http.StatusOK, "User deleted!")
}

