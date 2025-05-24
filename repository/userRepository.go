package repository

import (
	"posting-api/entity"
)

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
