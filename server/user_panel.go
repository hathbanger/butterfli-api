package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
)

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}
	
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.String(http.StatusOK, "Welcome "+name+"!")
}