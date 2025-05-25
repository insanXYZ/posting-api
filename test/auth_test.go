package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"posting-api/app"
	"posting-api/entity"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var server *echo.Echo
var db *gorm.DB

func init() {
	godotenv.Load("../.env", ".env")

	a := app.Init()
	server = a.GetEcho()
	db = a.GetDb()
}

func TestLogin(t *testing.T) {

	registerUser := entity.User{
		Username: faker.Name(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}

	b, _ := json.Marshal(registerUser)

	reqRegister := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewReader(b))
	reqRegister.Header.Set("Content-Type", "application/json")

	recRegister := httptest.NewRecorder()

	server.ServeHTTP(recRegister, reqRegister)
	assert.Equal(t, http.StatusOK, recRegister.Code)

	defer func() {
		deleteUser := &entity.User{
			Email:    registerUser.Email,
			Username: registerUser.Username,
		}

		err := db.Unscoped().Where(deleteUser).Delete(&entity.User{
			Email: registerUser.Email,
		}).Error

		assert.NoError(t, err)
	}()

	scenarios := []struct {
		name               string
		request            *entity.User
		expectedStatusCode int
	}{
		{
			name: "login success",
			request: &entity.User{
				Email:    registerUser.Email,
				Password: registerUser.Password,
			},
			expectedStatusCode: 200,
		},
		{
			name: "login failed (username or password wrong)",
			request: &entity.User{
				Email:    registerUser.Email,
				Password: "john doe password",
			},
			expectedStatusCode: 400,
		},
		{
			name: "login failed (error validation)",
			request: &entity.User{
				Email:    "loremipsumdolor",
				Password: "1234",
			},
			expectedStatusCode: 400,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			b, _ = json.Marshal(*scenario.request)

			reqLogin := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(b))
			reqLogin.Header.Set("Content-Type", "application/json")

			recLogin := httptest.NewRecorder()

			server.ServeHTTP(recLogin, reqLogin)

			assert.Equal(t, scenario.expectedStatusCode, recLogin.Code)
		})
	}
}
