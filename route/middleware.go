package route

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	HasJWT = echojwt.WithConfig(echojwt.Config{
		SuccessHandler: func(c echo.Context) {
			c.Set("user", c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims))
		},
		SigningKey:    []byte(os.Getenv("SECRET_KEY")),
		SigningMethod: echojwt.AlgorithmHS256,
	})
)
