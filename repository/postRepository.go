package repository

import (
	"context"
	"posting-api/entity"

	"gorm.io/gorm"
)

const (
	Limit_Pagination = 10
)

type PostRepository struct {
	Repository[entity.Post]
}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (p *PostRepository) TakeDetailPost(ctx context.Context, db *gorm.DB, entity *entity.Post) error {
	return db.WithContext(ctx).Joins("User").Preload("Liked").Preload("Comments").Find(entity).Error
}

func (p *PostRepository) TakeDetailPostsWithPagination(ctx context.Context, db *gorm.DB, entities *[]*entity.Post, page int) error {
	return db.WithContext(ctx).Joins("User").Preload("Liked").Preload("Comments").Offset(page * Limit_Pagination).Limit(Limit_Pagination).Find(entities).Error
}

func (p *PostRepository) Liked(ctx context.Context, db *gorm.DB, post *entity.Post, user *entity.User) error {
	return db.Model(post).Association("Liked").Append(user)
}

func (p *PostRepository) Unliked(ctx context.Context, db *gorm.DB, post *entity.Post, user *entity.User) error {
	return db.WithContext(ctx).Model(post).Association("Liked").Delete(user)
}

func (p *PostRepository) CountUserLikePost(ctx context.Context, db *gorm.DB, post *entity.Post, user *entity.User) int64 {
	return db.WithContext(ctx).Model(post).Where("users.id = ?", user.ID).Association("Liked").Count()
}

// func (p *PostRepository) TakeDetailPost(ctx context.Context, db *gorm.DB, entitiy *entity.Post) error {
// 	return db.WithContext(ctx).Joins("User")..Error
// }
