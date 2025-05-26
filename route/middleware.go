package route

import (
	"errors"
	"net/http"
	"os"
	"posting-api/dto/message"
	"posting-api/util"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	HasJWT, HasRefreshToken echo.MiddlewareFunc
)

func SetMiddleware() {
	HasJWT = echojwt.WithConfig(echojwt.Config{
		SuccessHandler: func(c echo.Context) {
			c.Set("user", c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims))
		},
		SigningKey:    []byte(os.Getenv("SECRET_KEY")),
		SigningMethod: echojwt.AlgorithmHS256,
	})

	HasRefreshToken = func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("refresh-token")
			if err != nil {
				return util.HttpResponseError(c, message.MISSING_REFRESH_TOKEN, nil, http.StatusUnauthorized)
			}

			if !strings.Contains(cookie.Value, "Bearer ") {
				return util.HttpResponseError(c, message.MISSING_REFRESH_TOKEN, nil, http.StatusUnauthorized)
			}

			tokenString := strings.Replace(cookie.Value, "Bearer ", "", -1)

			claims := jwt.MapClaims{}
			_, err = jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
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
				if _, ok := claims["exp"]; !ok {
					return util.HttpResponseError(c, message.INVALID_REFRESH_TOKEN, nil, http.StatusUnauthorized)
				}
			}

			return setAndNext()
		}
	}

}
