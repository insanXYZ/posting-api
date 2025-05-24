package route

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	HasJWT     echo.MiddlewareFunc
	RefreshJWT echo.MiddlewareFunc
)

func SetMiddleware() {
	HasJWT = echojwt.WithConfig(echojwt.Config{
		SuccessHandler: func(c echo.Context) {
			c.Set("user", c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims))
		},
		SigningKey:    []byte(os.Getenv("SECRET_KEY")),
		SigningMethod: echojwt.AlgorithmHS256,
	})

	RefreshJWT = func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				return c.NoContent(http.StatusUnauthorized)
			}

			if !strings.Contains(authorization, "Bearer ") {
				return c.NoContent(http.StatusUnauthorized)
			}

			tokenString := strings.Replace(authorization, "Bearer ", "", -1)
			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("invalid signing method")
				}

				return []byte(os.Getenv("SECRET_KEY")), nil
			})

			setAndNext := func() error {
				c.Set("user", jwt.MapClaims{
					"sub":  claims["sub"],
					"name": claims["name"],
				})

				return next(c)
			}

			if err != nil {
				if int64(claims["exp"].(float64)) <= time.Now().Unix() {
					return setAndNext()
				}

				return c.NoContent(http.StatusUnauthorized)
			}

			return setAndNext()
		}
	}
}
