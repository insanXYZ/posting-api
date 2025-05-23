package util

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreateJWT(claims jwt.MapClaims) (string, error) {
	newJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return newJwt.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func GetClaims(ctx echo.Context) jwt.MapClaims {
	return ctx.Get("user").(jwt.MapClaims)
}
