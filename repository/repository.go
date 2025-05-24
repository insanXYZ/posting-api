package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository[T any] struct{}

func (r *Repository[T]) Take(ctx context.Context, db *gorm.DB, entity *T) error {
	return db.WithContext(ctx).Take(entity).Error
}

func (r *Repository[T]) Create(ctx context.Context, db *gorm.DB, entity *T) error {
	return db.WithContext(ctx).Create(entity).Error
}

func (r *Repository[T]) Update(ctx context.Context, db *gorm.DB, entity *T) error {
	return db.WithContext(ctx).Model(entity).Updates(entity).Error
}

func (r *Repository[T]) Save(ctx context.Context, db *gorm.DB, entity *T) error {
	return db.WithContext(ctx).Save(entity).Error
}

func (r *Repository[T]) Delete(ctx context.Context, db *gorm.DB, entity *T) error {
	return db.WithContext(ctx).Delete(entity).Error
}
