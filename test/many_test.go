package test

import (
	"posting-api/config"
	"posting-api/entity"
	"testing"

	"github.com/joho/godotenv"
)

func TestInsertMany(t *testing.T) {
	godotenv.Load("../.env")
	db := config.NewGorm()

	user := &entity.User{
		ID: "4a94fe3d-e8ab-4931-b971-9d79c983dff1",
	}

	post := &entity.Post{
		ID: "70d7ebe2-a6ed-4a4e-b739-8df8b16a1293",
	}

	err := db.Model(post).Association("Liked").Append(user)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("ok")

}
