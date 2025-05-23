package repository

import (
	"context"
	"posting-api/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) TakeUserByEmail(ctx context.Context, db *gorm.DB, dst *entity.User, email string) error {
	return db.WithContext(ctx).Select("id", "username", "email", "password").Take(dst, "email = ?", email).Error
}

func (u *UserRepository) TakeUserById(ctx context.Context, db *gorm.DB, id string, dst *entity.User) error {
	return db.WithContext(ctx).Select("id", "username", "email").Where("id = ?", id).Take(dst).Error
}
