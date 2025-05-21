package main

import (
	"posting-api/app"

	"github.com/joho/godotenv"
)

func main() {
	// load environments
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	// running app
	err = app.Init().Run()
	if err != nil {
		panic(err.Error())
	}
}
