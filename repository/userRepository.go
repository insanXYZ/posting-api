package repository

import (
	"context"
	"posting-api/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) TakeUserByEmail(ctx context.Context, dst *entity.User, email string) error {
	return u.db.WithContext(ctx).Take(dst, "email = ?", email).Error
}

func (u *UserRepository) Create(ctx context.Context, value *entity.User) error {
	return u.db.Create(value).Error
}
