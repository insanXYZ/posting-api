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

func (p *PostRepository) PaginationPost(ctx context.Context, db *gorm.DB, dst *[]entity.Post, page int) error {
	return db.WithContext(ctx).Offset(page * Limit_Pagination).Limit(Limit_Pagination).Find(dst).Error
}
